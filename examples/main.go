package main

import (
	"random-user-generator/pkg/user"
)

func main() {
	randomUserData := user.NewRandomUserData("foobarbaz", 1, false, true)
	data, _ := randomUserData.Generate()
	randomUser := data[0]
	println(randomUser.Name.First)
	println(randomUser.Name.Last)
	println(randomUser.Location.Country)
	println(randomUser.Location.City)
	println(randomUser.Location.Street.Name)
	println(randomUser.Gender)
	println(randomUser.Nat)
	println(randomUser.Email)
}
