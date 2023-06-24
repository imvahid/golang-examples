package models

import "errors"

type Person struct {
	Name    string `json:"name"`
	Number  string `json:"number"`
	Address string `json:"address"`
}

var personList = []Person{}

func GetAllPersons() []Person {
	return personList
}

func GetPersonByName(name string) (*Person, error) {
	for _, person := range personList {
		if person.Name == name {
			return &person, nil
		}
	}
	return nil, errors.New("person not found")
}

func AddNewPersonByData(name, number, address string) (*Person, error) {
	person := Person{Name: name, Number: number, Address: address}
	personList = append(personList, person)
	return &person, nil
}

func AddNewPersonObject(p Person) {
	personList = append(personList, p)
}
