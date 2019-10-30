package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Id        float64 `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Address   string  `json:"address"`
	City      string  `json:"city"`
	Country   string  `json:"country"`
	Skill     string  `json:"skill"`
}

type filterfunc func(User, string) bool

func fromCountry(user User, country string) bool {
	return user.Country == country
}

func skillContains(user User, skill string) bool {
	return strings.Contains(strings.ToLower(user.Skill), strings.ToLower(skill))
}

func (users Users) filterList(fn filterfunc, value string, header string) {
	fmt.Println(header, value)
	for i := 0; i < len(users.Users); i++ {
		if fn(users.Users[i], value) {
			fmt.Println("User id#", users.Users[i].Id, "Name:", users.Users[i].FirstName, users.Users[i].LastName, "Address:", users.Users[i].Address, users.Users[i].City, users.Users[i].Country, "Skill:", users.Users[i].Skill)
		}
	}
	fmt.Println()
}

func main() {
	// Open our jsonFile
	filename := "../data/MOCK_DATA.json"
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened mock_data.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)
	fmt.Println("Unmarshaled users", len(users.Users))
	fmt.Println()

	users.filterList(fromCountry, "Australia", "Users from")
	// usersFilter(users, fromCountry, "China", "Users from")
	users.filterList(skillContains, "software", "Users with skill containing")
}
