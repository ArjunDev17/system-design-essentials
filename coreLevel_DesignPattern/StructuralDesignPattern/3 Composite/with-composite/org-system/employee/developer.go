package employee

import "fmt"

// Developer is a leaf node (no children)
type Developer struct {
	Name string
}

// ShowDetails prints developer info
func (d Developer) ShowDetails() {
	fmt.Println("Developer:", d.Name)
}