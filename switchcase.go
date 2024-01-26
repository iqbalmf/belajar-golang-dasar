package main

import "fmt"

func main() {
	name := "iqbal"
	switch name {
	case "iqbal":
		fmt.Println("Hello Iqbal")
	case "fauzan":
		fmt.Println("Hello Fauzan")
	default:
		fmt.Println("Hello, what's your name?")
	}

	//switch with short statement
	// length := len(name)
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	default:
		fmt.Println("Nama tidak ada")
	}
}
