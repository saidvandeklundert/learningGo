package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type PersonRegister struct {
	Persons []Person `json:"persons"`
}

func (p *PersonRegister) getPersonByName(name string) (Person, bool) {
	for _, person := range p.Persons {
		if person.Name == name {
			return person, true
		}
	}
	return Person{}, false
}

type PersonRepository interface {
	getPersonByName(name string) (Person, bool)
}

func worksWithRepository(pr PersonRepository) {
	person, exists := pr.getPersonByName("Jan")
	if exists {
		fmt.Println(person)
	}
	person, exists = pr.getPersonByName("DoesNotExist")
	if exists {
		fmt.Println(person)
	} else {
		fmt.Println("DoesNotExist")
	}

}
func main() {
	fmt.Println("The repository pattern")
	jsonData := []byte(`
	{
		"persons": 
			[
				{
					"name" : "Jan",
					"age" : 9 
				},
				{
					"name" : "Marie",
					"age" : 5
				}				
			]
	}

	`)
	var pr PersonRegister
	err := json.Unmarshal(jsonData, &pr)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Person register: ", pr)

	json_pr, _ := json.Marshal(pr)
	fmt.Println(string(json_pr))
	worksWithRepository(&pr)
}
