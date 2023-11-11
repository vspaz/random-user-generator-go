package main

import (
	"fmt"
	"random-user-generator/pkg/user"
)

func main() {
	randomUserData := user.NewRandomUserData("foobar", 1, true)
	data, _ := randomUserData.Generate()
	fmt.Println(data)
}
