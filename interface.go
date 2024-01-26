package main

import (
	"fmt"
	"math"
)

func main() {
	sayHellos(Person{Name: "Iqbal M Fauzan"})
	sayHellos(Animal{Name: "Anjing"})

	var kosong any = Ups()
	fmt.Println(kosong)

	fmt.Println("PERSEGI =====")
	bangunDatar := Persegi{15}
	fmt.Println("sisi: ", bangunDatar.sisi)
	fmt.Println("luas: ", bangunDatar.luas())
	fmt.Println("keliling: ", bangunDatar.keliling())

	fmt.Println("Lingkaran =====")
	lingkaran := Lingkaran{10}
	fmt.Println("diameter: ", lingkaran.diameter)
	fmt.Println("luas: ", lingkaran.luas())
	fmt.Println("keliling: ", lingkaran.keliling())
	fmt.Println("jari-jari: ", lingkaran.jarijari())

	fmt.Println("EMBEDDED INTERFACE")
	fmt.Println("BANGUNRUANG====")
	var bangunRuang hitung = &Kubus{10}
	kubusInstance := &Kubus{Sisi: 10}
	fmt.Println("sisi	   :", kubusInstance.sisiNya())
	fmt.Println("luas      :", bangunRuang.luas())
	fmt.Println("keliling  :", bangunRuang.keliling())
	fmt.Println("volume    :", bangunRuang.volume())


	fmt.Println("KOMBINASI SLICE & INTERFACE")
	//Dengan memanfaatkan slice dan interface{}, kita bisa membuat data array yang isinya adalah bisa apa saja.
	//data fruits yang terdiri dari Map, slice, dan string
	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"manggo", "pineapple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}
}

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func sayHellos(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

// interface kosong
func Ups() any {
	// return 1
	// return true
	return "Ups"
}

// case Bangun Datar & Lingkaran
type Hitung interface {
	luas() int
	keliling() int
}

type Lingkaran struct {
	diameter int
}

type Persegi struct {
	sisi float64
}

func (l Lingkaran) jarijari() float64 {
	return float64(l.diameter) / 2
}

func (linkaran Lingkaran) luas() float64 {
	return math.Pi * math.Pow(linkaran.jarijari(), 2)
}

func (lingkaran Lingkaran) keliling() float64 {
	return math.Pi * float64(lingkaran.diameter)
}

func (persegi Persegi) luas() float64 {
	return math.Pow(persegi.sisi, 2)
}

func (persegi Persegi) keliling() float64 {
	return persegi.sisi * 4
}

// embedded interface
type hitung2d interface {
	luas() float64
	keliling() float64
}
type hitung3d interface {
	volume() float64
}
type hitung interface {
	hitung2d
	hitung3d
}

type Kubus struct {
	Sisi float64
}

func (k *Kubus) volume() float64 {
	return math.Pow(k.Sisi, 3)
}
func (k *Kubus) luas() float64 {
	return math.Pow(k.Sisi, 2) * 6 //sisi^2 *6
}
func (k *Kubus) keliling() float64 {
	return k.Sisi * 12
}
func (k *Kubus) sisiNya() float64 {
	return k.Sisi
}
