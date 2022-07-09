// Maktabkhooneh
// The project of the 5th Season - the first part
package main

import "fmt"

type Car struct {
	name interface{}
	id   any
}

func (s *Car) SetterNameCar(name interface{}, id any) {
	s.name = name
	s.id = id
}
func main() {
	car := Car{}
	car.SetterNameCar(405, 1)
	fmt.Println(car.name)
	car.SetterNameCar("Paykan", "one")
	fmt.Println(car.name)
	car.SetterNameCar(map[int]string{131: "Prid", 132: "Prid", 141: "Prid"}, 478.25)
	fmt.Println(car.name, car.id)
}
