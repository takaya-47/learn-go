package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {
	var persons []Person
	for i := 0; i < 10_000_000; i++ {
		persons = append(persons, MakePerson("John", "Doe", 30))
	}
}
