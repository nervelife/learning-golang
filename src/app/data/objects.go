package data

import (
	"google.golang.org/appengine/datastore"
)

// AuthorEntity is for author entity in Datastore
type AuthorEntity struct {
	ID    int64
	Name  string `datastore:"name"`
	Alive bool   `datastore:"alive"`
}

// BookEntity is ...BookEntity
type BookEntity struct {
	Author   datastore.Key `datastore:"author"`
	Category string        `datastore:"category"`
	Title    string        `datastore:"title"`
}
