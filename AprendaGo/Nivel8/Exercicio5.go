package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (u byAge) Len() int {
	return len(u)
}

func (u byAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

func (u byAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

type byLastName []user

func (u byLastName) Len() int {
	return len(u)
}

func (u byLastName) Less(i, j int) bool {
	return u[i].Last < u[j].Last
}

func (u byLastName) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	for _, user := range users {
		sort.Strings(user.Sayings)
		for index, _ := range user.Sayings {
			user.Sayings[index] = "\n" + user.Sayings[index]
		}
	}

	fmt.Println(users)
	sort.Sort(byAge(users))
	fmt.Println(users)
	sort.Sort(byLastName(users))
	fmt.Println(users)

	// your code goes here

}

/*
Partindo do código abaixo, ordene os []user por idade e sobrenome.
    - https://play.golang.org/p/BVRZTdlUZ_
- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
*/
