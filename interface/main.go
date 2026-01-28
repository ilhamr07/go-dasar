package main

type Hasname interface {
	GetName() string
}

func sayHello(hasname Hasname) {
	println("Hello", hasname.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func main() { // This is a placeholder for the main function.
	person := Person{Name: "ilham mahli"}
	sayHello(person)
}
