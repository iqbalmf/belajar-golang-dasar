package main

import (
	"errors"
	"fmt"

)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian dengan NOL")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	// hasil, err := Pembagian(100, 0)
	// if err == nil {
	// 	fmt.Println("Hasil", hasil)
	// } else {
	// 	fmt.Println("Error", err.Error())
	// }

	/*
	* Custom Error
	 */
	err := SaveData("Iqbal", nil)
	if err != nil {
		//terjadi error
		// if validationErr, isOk := err.(*validationError); isOk {
		// 	fmt.Println("Validation Error:", validationErr.Error())
		// } else if notFoundError, isOk := err.(*notFoundError); isOk {
		// 	fmt.Println("Not Found Error:", notFoundError.Error())
		// } else {
		// 	fmt.Println("unknown Error", err.Error())
		// }

		switch finalError := err.(type){
		case *validationError:
			fmt.Println("Validation Error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Not Found Error:", finalError.Error())
		default:
			fmt.Println("unknown Error", err.Error())
		}
	} else {
		fmt.Println("Sukses Save Data")
	}
}

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.Message
}

type notFoundError struct {
	Message string
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}
	if id != "Iqbal" {
		return &notFoundError{"Data Not Found"}
	}
	return nil
}
