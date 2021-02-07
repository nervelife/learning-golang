package maxy

import "fmt"

// Maxy is a type
type Maxy struct {
	Planet string
	Size   int64
}

// CalcDistance measures distance from Earth
func (m *Maxy) CalcDistance() {
	fmt.Println("It's too far")
}