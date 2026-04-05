package employee

// Employee is the common interface (Component)
// Both Developer (leaf) and Manager (composite) will implement this
type Employee interface {
	ShowDetails()
}