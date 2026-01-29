package main

type Address struct {
	City, Province, Country string
}

func main() {

	address3 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address4 := &address3 //pointer

	address4.City = "Bandung"

	println(address3.City) //berubah menjadi bandung
	println(address4.City) //berubah menjadi bandung

}
