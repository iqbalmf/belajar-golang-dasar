package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal" // _ blank initializaiton
	"fmt"
)

func main()  {
	/* 
	* Initializations package
	*/
	fmt.Println(database.GetDatabase())	
}