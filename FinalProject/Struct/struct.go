package Struct

type People struct {
	User int `yaml:"user"`
	Person
}

type Person struct {
	Id      int    `yaml:"id"`
	Name    string `yaml:"name"`
	Age     int    `yaml:"age"`
	Address string `yaml:"address"`
}
