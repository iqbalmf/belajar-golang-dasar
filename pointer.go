package main

import "fmt"

type Address struct {
	City, Province, Country string
}
type Man struct {
	Name string
}

func main() {

	//pass by value
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1 //copy value

	address2.City = "Bandung"
	fmt.Println(address1) //tidak berubah
	fmt.Println(address2) //berubah menjadi bandung
	// cara ini duplikasi, membuat memory penuh tidak efektif

	//pass by reference (POINTER)
	var address3 *Address = &address1 //pointer
	address3.City = "Cianjur"
	fmt.Println("IKUT BERUBAH", address1)
	fmt.Println("POINTER", address3)

	/*
	* Pointer Address*/
	address4 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	// ChangeAddressToIndonesia(address4)//data country tidak berubah
	ChangeAddressToIndonesia(&address4) // pakai &, untuk mereference data Address
	fmt.Println(address4)

	/*
	*Pointer Method*/
	name := Man{"Iqbal"}
	name.Married()
	fmt.Println(name)
}

func (man *Man) Married() {
	//pakai *Man untuk mereference struct Man yang sama
	man.Name = "Mr. " + man.Name
}

func ChangeAddressToIndonesia(address *Address) { // tambah * untuk menjadikan reference
	address.Country = "Malaysia"
}
