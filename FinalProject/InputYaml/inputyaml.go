package Yaml

import (
	. "Yaml-v1/Struct"
)

func Initiate() []*People {
	Peoples := []*People{
		{
			User: 1,
			Person: Person{
				Id:      1,
				Name:    "saber khakbiz",
				Age:     28,
				Address: "\"Lahijan, Langaroud\""},
		},
		{
			User: 2,
			Person: Person{
				Id:      2,
				Name:    "Saman farajpour",
				Age:     30,
				Address: "\"Tehran, Mashhad\""},
		},
	}
	return Peoples

}
