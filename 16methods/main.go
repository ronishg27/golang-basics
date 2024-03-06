package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang;
	// struct is a custom data type;

	ronish := User{"Ronish", "ronish@go.dev", true, 20}
	// fmt.Println(ronish)
	// fmt.Println(ronish.Status)
	// fmt.Printf("Details of user %v is %+v\n", ronish.Name, ronish)

	// ronish.GetStatus()
	ronish.NewMail()
	fmt.Printf("Details of user %v is %+v\n", ronish.Name, ronish)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}
