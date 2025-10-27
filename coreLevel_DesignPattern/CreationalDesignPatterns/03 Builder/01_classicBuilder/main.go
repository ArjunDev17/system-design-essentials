package main

type Studen struct {
	Name  string
	Age   int
	Email string
	Phone string
}

type StudentBuilder struct {
	student *Studen
}

func NewStudentBuilder() *StudentBuilder {
	return &StudentBuilder{student: &Studen{}}
}

func (sb *StudentBuilder) SetName(name string) {
	sb.student.Name = name
}
func (sb *StudentBuilder) SetAge(age int) {
	sb.student.Age = age
}
func (sb *StudentBuilder) SetEmail(email string) {
	sb.student.Email = email
}
func (sb *StudentBuilder) SetPhone(phone string) {
	sb.student.Phone = phone
}
func (sb *StudentBuilder) Build() *Studen {
	return sb.student
}

func main() {
	sb := NewStudentBuilder()
	sb.SetName("arjunsingh")
	sb.SetAge(21)
	sb.SetEmail("arjun@gmail.com")
	sb.SetPhone("1234567890")

	student := sb.Build()
	println("Student Details:")
	println("Name:", student.Name)
	println("Age:", student.Age)
	println("Email:", student.Email)
	println("Phone:", student.Phone)

	println("Builder Pattern Example :", student)

}
