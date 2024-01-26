package main

import "fmt"

func main() {
	fmt.Println("DEFER")
	//fitur untuk memanggil sebuah function ketika eksekusi function selesai
	runApp()
}

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApp() {
	defer logging()
	fmt.Println("Run Applications")
}
