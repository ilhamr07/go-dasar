package main

import "fmt"

func random() any {
	return "gws bang ya"
}

func main() {

	var result any = random()
	fmt.Println(result)

	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown")
	}

}
