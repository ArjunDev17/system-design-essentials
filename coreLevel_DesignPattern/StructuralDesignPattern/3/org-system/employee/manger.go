package employee

import "fmt"

// Manager can only handle Developers ( problem)
type Manager struct {
    Name       string
    Developers []Developer
}

// AddDeveloper adds a developer to manager
func (m *Manager) AddDeveloper(dev Developer) {
    m.Developers = append(m.Developers, dev)
}

// ShowManagerDetails prints manager + team
func (m Manager) ShowManagerDetails() {
    fmt.Println("Manager:", m.Name)

    for _, dev := range m.Developers {
        dev.ShowDetails()
    }
}