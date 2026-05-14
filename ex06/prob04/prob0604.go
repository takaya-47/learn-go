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
	persons := make([]Person, 0, 10_000_000) // 事前にcapacityを確保。中身は空（lengthが0）。
	for i := 0; i < 10_000_000; i++ {
		persons = append(persons, MakePerson("John", "Doe", 30))
	}
}
