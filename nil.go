package main

import "fmt"

func main() {
	data:= NewMap("")
	if data == nil {
		fmt.Println("data map masih kosong")
	}else {
		fmt.Println(data["name"])
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

// func ContohNilTidakBisa(name string) string {
// 	if name == "" {
// 		return nil 
// 	} else {
// 		return name
// 	}
// }
// NIL HANYA BISA DIGUNAKAN UNTUK TIPE DATA TERTENTU (MAP, SLICE, INTERFACE, FUNCTION, POINTER, CHANNEL)