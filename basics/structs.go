package main

import "fmt"

type Address struct {
	addressLine string
	zipCode     string
	city        string
	country     string
}

type Person struct {
	firstName string
	lastName  string
	address   Address
}

func (p *Person) updateAddress(addr *Address) {
	p.address = *addr
}

func (p *Person) print() {
	fmt.Printf("%+v\n", p)
}

func runStructsExample() {
	johnDoe := Person{
		firstName: "John",
		lastName:  "Doe",
		address: Address{
			addressLine: "21 2nd St.",
			zipCode:     "07302",
			city:        "Jersey City",
			country:     "US",
		},
	}
	johnDoe.print()

	johnDoe.updateAddress(&Address{
		addressLine: "287 10th Ave",
		zipCode:     "10001",
		city:        "New York",
		country:     "US",
	})
	johnDoe.print()
}
