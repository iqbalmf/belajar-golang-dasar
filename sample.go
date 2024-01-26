package main

import "fmt"

func main() {
	fmt.Println("This is sample main")

	//data type
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga koma lima = ", 3.5)

	//boolean
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	//function string
	fmt.Println(len("Iqbal M Fauzan"))
	fmt.Println("Iqbal M Fauzan"[5])

	//variable
	fmt.Println("VARIABLE==")
	// var name = ""
	name := ""
	name = "Iqbal M Fauzan"
	fmt.Println(name)
	name = "Iqbal"
	fmt.Println(name)

	var (
		firstName = "Iqbal"
		lastName  = "Fauzan"
	)
	fmt.Println(firstName, lastName)

	//constant
	const constFirstName string = "Iqbal"
	const constLastName = "Fauzan"
	fmt.Println("constant", constFirstName, constLastName)

	const (
		fsName string = "iqbal"
		lsName        = "Fauzan"
	)
	fmt.Println("multiple constant", fsName, lsName)

	//konversi tipe data
	fmt.Println("Konversi Tipe Data")
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) //number overflow
	fmt.Println(nilai32, nilai64, nilai16)

	var namee = "Iqbal Fauzan"
	var e uint8 = namee[6]
	var eString = string(e)
	fmt.Println(eString)

	//type declarations
	fmt.Println("Type Declarations")
	type NoKTp string

	var ktpIqbal NoKTp = "11111"
	var contoh string = "22222"
	var contohKtp NoKTp = NoKTp(contoh)
	fmt.Println(ktpIqbal)
	fmt.Println(contohKtp)
}
