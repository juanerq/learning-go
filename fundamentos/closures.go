package main

import "fmt"

type User struct {
	Id   int
	Name string
}

func CacheUser() (func(id int) *User, func(user User)) {
	var users []User

	return func(id int) *User {
			for i := range users {
				if users[i].Id == id {
					return &users[i]
				}
			}
			return nil
		},
		func(user User) {
			users = append(users, user)
		}

}

func main() {
	getUser, addUser := CacheUser()

	addUser(User{Id: 1, Name: "Juan"})
	user := getUser(1)
	user.Name = "JuanERQ"

	fmt.Println(getUser(1))

}
