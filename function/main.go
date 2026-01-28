package main

import "fmt"

//func tanpa parameter dan tanpa return value
func sayHello() {
	fmt.Println("Hello, World!")
}

//func parameter
func sayHelloParameter(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

//func return value
func getHello(name string) string {
	return "Hello " + name
}

//fungsi multiple return value
func getFullName() (string, string) {
	return "ilham", "mahli"
}

//fungsi menghiraukan return value

//named return value
func getCompleteName() (fullName, middleName, lastName string) {
	fullName = "ilham"
	middleName = "mahli"
	lastName = "ganteng"

	return fullName, middleName, lastName
}

//varadiac function
func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

//function as value

func getGoodBye(name string) string {
	return "Goodbye " + name
}

//function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "bangsat" {
		return "****"
	} else {
		return name
	}
}

//anonymous function

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blacklisted", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//defer panic recover function

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging()

	fmt.Println("Run Application")

}

func runApp(error bool) {
	defer logging()

}

func main() {
	sayHello()
	sayHelloParameter("ilham", "mahli")
	fmt.Println(getHello("mahli"))

	firstName, _ := getFullName()
	fmt.Println(firstName)

	a, b, c := getCompleteName()
	fmt.Println(a, b)
	fmt.Println(b, c)
	fmt.Println(c, a)

	fmt.Println(sumAll(10, 10, 10, 10, 10))

	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("ilham"))
	fmt.Println(misal("mahli"))

	sayHelloWithFilter("mahli", spamFilter)

	filterNama := spamFilter
	sayHelloWithFilter("bangsat", filterNama)

	blacklist := func(name string) bool {
		return name == "bangsat"
	}
	registerUser("ilham", blacklist)

	//anonymous function
	registerUser("bangsat", func(name string) bool {
		return name == "bangsat"
	})

	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursif(10))

	runApplication()
}

//recursif function vs loop function
//loop function
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

//recursif function vs loop function
//recursif function

func factorialRecursif(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursif(value-1)
	}
}
