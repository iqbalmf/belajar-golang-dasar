package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main()  {
	result := helper.SayHello("Iqbal")
	fmt.Println(result)

	fmt.Println(helper.Applications)
	// fmt.Println(helper.version) // tidak bisa akses
	// fmt.Println(helper.sayGoodBye("Iqbal")) //tidak bisa akses
}