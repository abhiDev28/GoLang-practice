package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p person) updateNameByValue(newFirstName string) person {
	p.firstName = newFirstName
	return p
}

func (p *person) updateNameByRef(newFirstName string) {
	(*p).firstName = newFirstName
}
