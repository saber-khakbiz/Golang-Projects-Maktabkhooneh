package Generate

import (
	. "Yaml-v1/InputYaml"
	"fmt"
	"io/ioutil"
	"log"
)

func YamlFile() {
	peoples := Initiate()
	data := []byte(fmt.Sprintf("---\r\n"))
	data = append(data, []byte(fmt.Sprintf("People:\n"))...)
	for _, people := range peoples {
		data = append(data, []byte(fmt.Sprintf("  User %v:\n", people.User))...)
		data = append(data, []byte(fmt.Sprintf("    id: %v\n", people.Person.Id))...)
		data = append(data, []byte(fmt.Sprintf("    Name: %v\n", people.Person.Name))...)
		data = append(data, []byte(fmt.Sprintf("    Age: %v\n", people.Person.Age))...)
		data = append(data, []byte(fmt.Sprintf("    Address: %v\n", people.Person.Address))...)
	}

	data = append(data, []byte(fmt.Sprintf("..."))...)

	err := ioutil.WriteFile("out.yml", data, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
