package main

import "fmt"

func main() {
	var customer Customer
	fmt.Println(customer)

	customer.name = "Iqbal Fauzan"
	customer.addresss = "Indonesia"
	customer.age = 27
	fmt.Println(customer, customer.name)

	//struct literals
	customerNew := Customer{"IqbalNew", "Indonesia New", 27}
	fmt.Println(customerNew)

	//struct method
	iqbal:= Customer{name: "Iqbal"}
	iqbal.sayHello("Joko")
}

type Customer struct {
	name, addresss string
	age            int
}

//struct method
func (customer Customer) sayHello(name string){
	fmt.Println("Hello", name, "my name is", customer.name)
}
