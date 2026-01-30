package main

type Address struct {
	City, Province, Country string
}

func main() {
	//pass by value
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := address1

	address2.City = "Bandung"

	println(address1.City) //tidak berubah
	println(address2.City) //berubah menjadi bandung

	//pass by reference memory

	//pass by value
	address3 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address4 := &address3 //pointer

	address4.City = "Bandung"

	println(address3.City) //berubah menjadi bandung
	println(address4.City) //berubah menjadi bandung

}
