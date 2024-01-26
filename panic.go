package main

import "fmt"

func main() {
	fmt.Println("PANIC")
	//fitur untuk menghentikan program
	//panic function dipanggil, defer tetap akan dieksekusi
	runApps(true)
}

func endApp() {
	fmt.Println("End App")
}

func runApps(errors bool) {
	defer endApp()
	if errors {
		panic("Ups Error bos")
	}
}
