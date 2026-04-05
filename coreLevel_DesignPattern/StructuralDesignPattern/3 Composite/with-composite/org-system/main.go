package main

import "org-system/employee"

func main() {

	// Create leaf nodes (Developers)
	dev1 := employee.Developer{Name: "Arjun"}
	dev2 := employee.Developer{Name: "Rahul"}
	dev3 := employee.Developer{Name: "Neha"}

	// Create manager1 and add developers
	manager1 := &employee.Manager{Name: "Amit"}
	manager1.Add(dev1)
	manager1.Add(dev2)

	// Create CEO (top-level manager)
	ceo := &employee.Manager{Name: "Rohit"}

	// Add manager1 under CEO (THIS was impossible before  now possible ✅)
	ceo.Add(manager1)

	// Add direct developer under CEO
	ceo.Add(dev3)

	// Single uniform call
	ceo.ShowDetails()
}