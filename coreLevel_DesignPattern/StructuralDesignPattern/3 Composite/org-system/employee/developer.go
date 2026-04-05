package employee

import "fmt"

// Developer is a simple leaf node
type Developer struct {
    Name string
}

// ShowDetails prints developer info
func (d Developer) ShowDetails() {
    fmt.Println("Developer:", d.Name)
}