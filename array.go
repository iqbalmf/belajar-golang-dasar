package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Iqbal"
	names[1] = "M"
	names[2] = "Fauzan"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		10, 20, 1,
	}
	fmt.Println(values)

	//function array
	fmt.Println("panjang array", len(values))
	var valuesArray = [...]int{
		10, 20, 30, 40,
	}
	fmt.Println(valuesArray)
	fmt.Println(valuesArray[2])
	fmt.Println(len(valuesArray))
}
