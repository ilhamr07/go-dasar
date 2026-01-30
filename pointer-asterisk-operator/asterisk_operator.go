package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address3 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address4 := &address3 //pointer

	address4.City = "Bandung"

	fmt.Println(address4) //berubah menjadi bandung

	//pointer asterisk operator
	*address4 = Address{"Surabaya", "Jawa Timur", "Indonesia"}

	fmt.Println(address3)
	fmt.Println(address4)

}
