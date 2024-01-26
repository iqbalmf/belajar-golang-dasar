package main

import "fmt"

func main() {
	// var person map[string]string = map[string]string{}
	person := map[string]string{
		"name":    "iqbal",
		"address": "Cianjur",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	//Function MAP
	// book := map[string]string
	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Iqbal"
	book["Ups"] = "Salah"
	fmt.Println(book)
	delete(book, "Ups")
	fmt.Println(book)
}
