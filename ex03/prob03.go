package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	employee1 := Employee{"太郎", "山田", 20}
	employee2 := Employee{
		firstName: "花子",
		lastName:  "田中",
		id:        30,
	}
	var employee3 Employee
	employee3.firstName = "次郎"
	employee3.lastName = "鈴木"
	employee3.id = 40

	fmt.Println(employee1)
	fmt.Println(employee2)
	fmt.Println(employee3)
}
