package main

import "fmt"

func main() {
	abhi := person{
		firstName: "Abhi",
		lastName:  "Nilange",
		contact: contactInfo{
			email:   "abhi@gmail.com",
			zipCode: "400211",
		},
	}

	updatedAbhi := abhi.updateNameByValue("Abhishek")
	fmt.Printf("**** Updated Pass by value ****")
	updatedAbhi.print()
	fmt.Printf("**** Updated Pass by rceference ****")
	personPointer := &abhi
	personPointer.updateNameByRef("Abhishek")
	abhi.print()
}
