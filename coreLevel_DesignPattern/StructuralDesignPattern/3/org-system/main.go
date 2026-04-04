package main

import "org-system/employee"

func main() {

    // Create developers
    dev1 := employee.Developer{Name: "Arjun"}
    dev2 := employee.Developer{Name: "Rahul"}
    dev3 := employee.Developer{Name: "Neha"}

    // Create manager
    manager1 := employee.Manager{Name: "Amit"}
    manager1.AddDeveloper(dev1)
    manager1.AddDeveloper(dev2)

    // Try to create CEO
    ceo := employee.Manager{Name: "Rohit"}

    //  PROBLEM: Cannot add manager1 under CEO
    // ceo.AddDeveloper(manager1) // compile error

    // Workaround (bad)
    ceo.AddDeveloper(dev3)

    // Output
    manager1.ShowManagerDetails()
    ceo.ShowManagerDetails()
}