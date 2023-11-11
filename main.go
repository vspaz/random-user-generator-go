package main

import (
	"fmt"
	"random-user-generator/pkg/user"
)

func main() {
	randomUserData := user.NewRandomUserData("foobarbaz", 1, false, true)
	data, _ := randomUserData.Generate()
	fmt.Println(data)
}
