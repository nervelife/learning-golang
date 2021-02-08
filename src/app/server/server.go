package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/datastore"
	"github.com/nervelife/learning-golang/src/app/data"
	"google.golang.org/api/iterator"
)

var ctx context.Context
var projectID string
var client datastore.Client

func init() {

	fmt.Println("Initializing server...")
	ctx = context.Background()

	projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Fatal("Project ID Not Found")
	}

	// GOOGLE_APPLICATION_CREDENTIALS
	c, err := datastore.NewClient(ctx, projectID)

	if err != nil {
		log.Fatalf("Failed to create client %v", err)
	}

	client = *c

	log.Println("Initilized...")
}

// Run the server
func Run() {

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/save-author", SaveAuthorHandler)
	http.HandleFunc("/get-all-authors", GetAllAuthors)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

// IndexHandler is an handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, World!")
}

// SaveAuthorHandler is an handler
func SaveAuthorHandler(w http.ResponseWriter, r *http.Request) {
	a := data.AuthorEntity{
		Name:  "Giang giang",
		Alive: true,
	}
	aKey := datastore.IncompleteKey("Authors", nil)
	if _, err := client.Put(ctx, aKey, &a); err != nil {
		log.Fatalf("Error saving to datastore %v", err)
	}
}

// GetAllAuthors is an function
func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	query := datastore.NewQuery("Authors").Filter("alive >", false)

	it := client.Run(ctx, query)

	var authors []data.AuthorEntity

	for {
		var author data.AuthorEntity
		key, err := it.Next(&author)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Error fetching next author: %v", err)
		}
		author.ID = key.ID
		authors = append(authors, author)
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}
