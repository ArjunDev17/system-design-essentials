package employee

import "fmt"

// Manager is a composite node (can contain other employees)
type Manager struct {
	Name      string
	Employees []Employee // IMPORTANT: interface type (key of composite)
}

// Add adds any Employee (Developer or Manager)
func (m *Manager) Add(emp Employee) {
	m.Employees = append(m.Employees, emp)
}

// ShowDetails prints manager + recursively all children
func (m Manager) ShowDetails() {
	fmt.Println("Manager:", m.Name)

	// Recursive call
	for _, emp := range m.Employees {
		emp.ShowDetails()
	}
}