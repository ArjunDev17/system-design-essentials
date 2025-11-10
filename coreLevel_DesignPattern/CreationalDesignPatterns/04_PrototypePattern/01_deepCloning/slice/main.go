package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Marks []string
}

func main() {
	s1 := Student{
		Name:  "Kabbu",
		Marks: []string{"99", "98", "100"},
	}
	fmt.Printf("%+v :", s1)
	s2 := Student{
		Name:  "Arjun",
		Marks: make([]string, len(s1.Marks)),
	}
	copy(s2.Marks, s1.Marks) //source and destination
	// copy(s1.Marks, s2.Marks)
	fmt.Printf("%+v :", s2)
}
