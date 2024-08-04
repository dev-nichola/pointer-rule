package main

import "fmt"

func main() {
	var a int

	a = 10
	fmt.Println(a)

	fmt.Println(student().Name)

	i := pointer()

	fmt.Println(*i)

	var num int32 = 32

	cetakNumber(int(num))

	var std Student = Student{
		Name: "Nichola Joko Morro",
	}
	printStudent(&std)
}

// Tipe Data dengan Value
// Tipe Data dengan Pointer

type Student struct {
	Id   int
	Name string
	Age  int
}

func student() Student {
	student := Student{}

	student.Id = 1
	student.Name = "Nichola Saputra"
	student.Age = 20

	return student
}

func pointer() *int {
	var i int

	i = 100_000_000

	return &i
}

func cetakNumber(num int) {
	fmt.Println(num)
}

func printStudent(std *Student) {
	fmt.Println(std.Name)
}
