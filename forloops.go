package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke-", counter)
		counter++
	}
	fmt.Println("selesai")

	//for dengan statement
	for count := 1; count <= 10; count++ {
		fmt.Println("Perulangan ke-", count)
	}

	//for range
	names := []string{"Iqbal", "Muhammad", "Fauzan"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	//break & continue
	fmt.Println("Brak & Continue")
	fmt.Println("Break")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Perulangan ke-", i)
	}
	fmt.Println("Continue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Perulangan ke-", i)
	}
}
