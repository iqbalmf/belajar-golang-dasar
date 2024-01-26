package main

import "fmt"

func main() {
	name := "Iqbal"

	if name == "Iqbal" {
		fmt.Println("Hi, ", name)
	} else {
		fmt.Println("Hi, Nama kamu siapa?")
	}

	//if short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
