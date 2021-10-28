package main

import (
	"fmt"
	"os"
)

/*
var (
	name, course         string // nicht definierte Variablen werden mit 0 oder ''
	module               float64
	street, number, city = "Unter den Linden", 4, "Berlin"
)
*/
func main() {

	var name string
	num_int := 5
	num_fl := 6.5

	result := float64(num_int) + num_fl
	ptr := &result

	name = "Sascha"

	//fmt.Println("Hello from", runtime.GOOS) // runtime operating system
	fmt.Println("Result: ", &result, " PTR ", *ptr)
	fmt.Println("Namee: ", name)

	changeResult(&result)
	name = changeName(name)

	fmt.Println("Result: ", &result, " PTR ", *ptr)
	fmt.Println("Namee: ", name)

	states := []string{"Alabame", "Utah", "Nebraska"}

	for _, i := range states {
		fmt.Println(i)
	}

	fmt.Println("Aenderungen für den Github Dummy")
}

func changeResult(result *float64) {
	*result = 56.6
}

func changeName(namee string) string {

	namee = os.Getenv("USERNAME")

	return namee
}
