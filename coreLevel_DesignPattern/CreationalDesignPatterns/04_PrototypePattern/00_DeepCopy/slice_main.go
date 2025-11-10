package main

import "fmt"

// type slice struct{
// 	// ptr *T//poniter to the the underline data structure array
// 	len int
// 	cap int
// }

type Student struct {
	Name    string
	Subject []string
}

func main() {
	stu := Student{
		Name:    "Arjun",
		Subject: []string{"Math", "com", "Science"},
	}
	stuShalloC := stu

	stuShalloC.Subject[0] = "drawing"
	fmt.Println("orignal student :", stu)

	fmt.Println("Shallow copy student :", stuShalloC)

	//deep Cloning
	stu_deep1 := Student{
		Name:    "Ram",
		Subject: []string{"Atharved", "Samved"},
	}
	fmt.Println("s :", stu_deep1)

}

/*


/*
* list of employees {{
* id
* name
* salary
* department
*
* java 8 stream API name by assending order and then by salry by desending
* */

// void main() {
//       Employee emp = new Employee(1, 122.40,"IT");
//       Employee emp2 = new Employee(2, 11.40,"admin");
//       Employee emp3= new Employee(3, 19.40,"can");
//     List<Employee> employees = new ArrayList<>();
//     employees.add(emp);
//     employees.add(emp2);
//     employees.add(emp3);
//     //
//     List<Employee> result=  employees.stream().sorted((e)->e.salary()).thenComparing(Employee::department).reversed()).collect(Collectors.toList());
//     System.out.println(result);

// }
