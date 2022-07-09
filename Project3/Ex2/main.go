// Maktabkhooneh
// The project of the third season - the Second part

package main

import "fmt"

// Person structure
type Person struct {
	Firstname string
	Lastname  string
	Address   string
	Age       int
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
	var persons []*Person

	// Define first person
	p1 := &Person{}
	p1.SetterPerson(
		"Saber",
		"Khakbiz",
		"Lahijan-Gilan-Iran",
		28)
	persons = append(persons, p1)

	// Define Second person
	p2 := &Person{}
	p2.SetterPerson(
		"Ramin",
		"Farajpour",
		"Tehran-Iran",
		35)
	persons = append(persons, p2)

	//Show
	for _, person := range persons {
		fmt.Printf("Name: %s, Family: %s, Addr: %s, Age: %v \n",
			person.Firstname,
			person.Lastname,
			person.Address,
			person.Age)
	}

}
