package main

import "fmt"

func main() {
	sayHello()
	sayHelloTo("Iqbal", "Fauzan")
	result := returnValue("Iqbal")
	fmt.Println(result)

	firstName, lastName, _ := getFullName()
	fmt.Println(firstName, lastName)

	total := sumAll(10, 9, 8, 7, 6)
	fmt.Println(total)
	numbers := []int{10, 9, 8, 7, 6}
	fmt.Println(sumAll(numbers...)) //diubah ke array dari slice

	goodbye := getGoodBye //getGoodBye() => function, getGoodBye variable
	fmt.Println(goodbye("Iqbal"))

	sayHelloWithFilter("Iqbal", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	blacklist := func(name string) bool {
		return name == "anjing"
	}
	registerUser("iqbal", blacklist)
	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	})

	factorialLoop := factorialLoop(10)
	factorialRecursive := factorialRecursive(10)
	fmt.Println(factorialLoop, factorialRecursive)
}
func sayHello() {
	fmt.Println("Hello")
}

// function parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// function return value
func returnValue(name string) string {
	return "Hello " + name
}

// function return multiple values
func getFullName() (string, string, int) {
	return "Iqbal", "Fauzan", 1
}

// function varargs
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// function value
func getGoodBye(name string) string {
	return "Good Bye " + name
}

// function as parameter
func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello " + filter(name))
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

// function type declarations
type Filter func(string) string

func sayHelloWithFilterTypeDeclarations(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello, " + filteredName)
}

// anonymouse function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// recursive function
func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
