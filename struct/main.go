package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
	Married bool
}

//struct method
func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	// This is a placeholder for the main function.
	var ilham Customer
	ilham.Name = "ilham mahli"
	ilham.Address = "jl. merdeka no. 10"
	ilham.Age = 20
	ilham.Married = false

	fmt.Println(ilham)

	mahli := Customer{
		Name:    "mahli ilham",
		Address: "jl. sudirman no. 20",
		Age:     21,
		Married: true,
	}
	fmt.Println(mahli)

	mahli.sayHello("agus")
}
