package main

import "fmt"

func main() {
	// var slice [*tidak ada initiate jumlah isi*]string
	names := [...]string{"iqbal", "fauzan", "joko", "budi", "nugraha"} // array
	slice1 := names[2:5]                                               //slice
	fmt.Println("slice 1 ", slice1)

	slice2 := names[:3]
	fmt.Println("slice 2 ", slice2)
	slice3 := names[3:]
	fmt.Println("slice 3 ", slice3)
	slice4 := names[:]
	fmt.Println("slice 4 ", slice4)

	//function slice

	fmt.Println("\n")
	days := [...]string{
		"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu",
	}
	dayslice1 := days[5:]
	fmt.Println(dayslice1)
	dayslice1[0] = "Sabtu Baru"
	dayslice1[1] = "Minggu Baru"
	fmt.Println(dayslice1)
	fmt.Println(days)

	//APPEND
	fmt.Println("APPEND===")
	dayslice2 := append(dayslice1, "Libur Baru")
	dayslice2[0] = "Sabtu Lama"
	fmt.Println(dayslice2)
	fmt.Println(days)

	//MAKE
	fmt.Println("MAKE===")
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Iqbal"
	newSlice[1] = "Fauzan"
	// newSlice[2] = "test" => error out of range
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	newSlice2 := append(newSlice, "GOLANG")
	fmt.Println(newSlice2, len(newSlice2), cap(newSlice2))
	//COPY
	fmt.Println("COPY===")
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//ARRAY vs SLICE
	fmt.Println("ARRAY VS SLICE")
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
