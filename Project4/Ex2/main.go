// Maktabkhooneh
// The project of the Fourth Season - the Second part
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name   any
	Family any
	Age    any
	City   any
}

func map2struct(m map[any]any) Person {
	s := Person{}
	s.Name = m["Name"]
	s.Family = m["Family"]
	s.Age = m["Age"]
	s.City = m["City"]
	return s
}

func main() {
	m := map[any]any{
		"Name":   "Saber",
		"Family": "Khakbiz",
		"Age":    28,
		"City":   "Lahijan",
	}
	fmt.Println(m, "-->", reflect.TypeOf(m))
	s := map2struct(m)
	fmt.Printf("{Name:%v, Family:%v, Age:%v, City: %v}--> ",
		s.Name,
		s.Family,
		s.Age,
		s.City,
	)
	fmt.Println(reflect.TypeOf(s))
}
