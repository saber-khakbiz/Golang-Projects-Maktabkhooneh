// Maktabkhooneh
// The project of the third season - the first part

package main

import "fmt"

// Person structure
type Person struct {
	Firstname string
	Lastname  string
	Address   string
	Age       int
}

// GetterPerson Methode
func (p Person) GetterPerson() Person {
	return p
}

// SetterPerson Methode
func (p *Person) SetterPerson(
	firstname string,
	lastname string,
	addr string,
	age int) {
	p.Firstname = firstname
	p.Lastname = lastname
	p.Address = addr
	p.Age = age
}

// Main function
func main() {
	p1 := Person{}
	p1.SetterPerson(
		"saber",
		"khakbiz",
		"Lahijan-Gilan-Iran",
		28)
	fmt.Println(p1.GetterPerson())

}
