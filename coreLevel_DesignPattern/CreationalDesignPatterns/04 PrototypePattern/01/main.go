package main

import "fmt"

//if you are going to impliment prototype remember it is a kinf of interface which will return always prptptype having clone method

type Prototype interface {
	Clone() Prototype
}

// for java developer struct is a kind of class where we will imliment Prototype interface and here not required or given any keyword where we we will use like in java
// exted or impliment
type Person struct {
	Name       string
	Age        int
	Profession string
	HObby      []string
}

func (p *Person) Clone() Prototype {
	newP := *p
	return &newP
}
func main() {
	//Create an orignal Prototype
	orignal := Person{
		Name:  "ArjunSingh",
		Age:   28,
		HObby: []string{"Coding", "Badminton"},
	}
	clone := orignal.Clone().(*Person)
	//modefy clone fields
	clone.Name = "Kabbu"
	clone.Age = 4
	// clone.HObby = []string{"dancing", "walking"}
	clone.HObby[0] = "dancing"
	fmt.Println("Orignal :", orignal)
	fmt.Println("Clone :", clone)
}

//issue

// arjunsingh@surajsingh 01 % go run main.go
// Orignal : {ArjunSingh 28  [dancing Badminton]}
// Clone : &{Kabbu 4  [dancing Badminton]}
// arjunsingh@surajsingh 01 %
