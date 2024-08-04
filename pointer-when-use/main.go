package main

import (
	"fmt"
)

type Student struct {
	Id      int
	Name    string
	Age     int
	Address *string
}

func (student *Student) setName(name string) {
	student.Name = name
}

func setStudentName(student *Student, name string) {
	student.Name = name
}

func main() {
	var request Student = Student{
		Id:   123,
		Name: "Nichola",
		Age:  20,
	}

	fmt.Println(request.Address)

	// // mengecek size struct
	// fmt.Println(unsafe.Sizeof(student))
	// fmt.Println(unsafe.Sizeof(student.Id))
	// fmt.Println(unsafe.Sizeof(student.Name))
	// fmt.Println(unsafe.Sizeof(student.Age))

	// fmt.Println("Before name : ", student.Name)

	// student.setName("Saputra")

	// fmt.Println("Afer name : ", student.Name)
}
