package main

import "fmt"

func main() {
	fmt.Println("RECOVER")
	//function digunakan untuk menangkan data panic
	runAppr(true)

}

func runAppr(errors bool) {
	defer endApp()
	if errors {
		panic("ERROR")
	}
}

func endApp() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Error", message)
}

//penanganan error aplikasi seperti try/catch di bahasa lain