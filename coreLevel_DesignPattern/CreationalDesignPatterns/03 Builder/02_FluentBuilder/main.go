package main

type Student struct {
	Name       string
	rollNumber string
}

// setone builder
type StudentBuilder struct {
	student *Student
}

func NewStudentBuilder() *StudentBuilder {
	return &StudentBuilder{student: &Student{}}
}

func (sb *StudentBuilder) SetName(name string) *StudentBuilder {
	sb.student.Name = name
	return sb

}
func (sb *StudentBuilder) SetRoll(rn string) *StudentBuilder {
	sb.student.rollNumber = rn
	return sb
}

func main() {
	println("Hello world")

	res := NewStudentBuilder().SetName("Ram").SetRoll("234")
	println(res.student.Name)
	println(res.student.rollNumber)
}
