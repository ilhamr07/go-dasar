package main

import "fmt"

type Customer struct {
	Name    string
	Address string
	Age     int
	Married bool
}

func main() {
	// This is a placeholder for the main function.
	var ilham Customer
	ilham.Name = "ilham mahli"
	ilham.Address = "jl. merdeka no. 10"
	ilham.Age = 20
	ilham.Married = false

	fmt.Println(ilham)
}
