package main

import (
	"random-user-generator/pkg/user"
)

func main() {
	randomUserData := user.NewRandomUserData("foobarbaz", 1, false, true)
	data, _ := randomUserData.Generate()
	userName := data[0].Name
	print(userName.First)
	print(userName.Last)
}
