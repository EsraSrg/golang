package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name           string    `json:"name"`
	Surname        string    `json:"surname"`
	Age            int       `json:"age"`
	IsMale         bool      `json:"isMale"`
	Height         float32   `json:"ppp"`
	FavoriteSports []string  `json:"favoriteSports"`
	SubUsers       []SubUser `json:"subUsers"`
	X              interface{}
}

type SubUser struct {
	Name string `json:"name"`
}

var jsonObject string = `{
	"name": "Osman",
	"surname": "AYDUR",
	"age": 250,
	"isMale": true,
	"ppp": 193,
	"favoriteSports": [
		"basketball",
		"volleyball"
	],
	"subUsers": [
		{
			"name": "Esra"
		},
		{
			"name": "Akif"
		}
	],
	"x":[]
}`

func main() {
	// var user User

	// user.Name = "Deniz"

	// a, b := json.Marshal(user)
	// if b != nil {
	// 	fmt.Println("Error:", b.Error())
	// } else {
	// 	fmt.Println(string(a))
	// }

	var user User
	c := json.Unmarshal([]byte(jsonObject), &user)
	if c != nil {
		fmt.Println("Error:", c.Error())
	}
	fmt.Println(reflect.TypeOf(user.X))
}

/*
Camel case
osmanAydur
firatUniversitesiYazilimMuhendisligi

Pascal Case
OsmanAydur
FiratUniversitesiYazilimMuhendisligi
*/
