package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string `json:"name"`
	Surname string
	Email   string
	Phone   string
	Age     string
}

type Book struct {
	Name        string
	Description string
	Author      string
	ISBM        string
}

type Role struct {
	Name        string
	Description string
}

// Kullanıcıyı veritabanına kayıt eder.
func (u *User) Create() {}
func (u *User) Get()    {}
func (u *User) List()   {}
func (u *User) Update() {}
func (u *User) Delete() {}

func main() {
	var user User
	user.Age = "10"
	user.Create()

	response, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Hata var: ", err.Error())
	} else {
		fmt.Println(string(response))
	}
}
