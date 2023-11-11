package main

import (
	"fmt"
	"random-user-generator/pkg/user"
)

func main() {
	randomUserData := user.NewRandomUserData("foobar", 1, false)
	data, _ := randomUserData.Generate()
	fmt.Println(data[0].Name.First)
}
