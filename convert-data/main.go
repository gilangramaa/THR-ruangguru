package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	userData := strings.Split(data, ",")
	age, _ := strconv.Atoi(userData[1])

	return User{
		Name:    userData[0],
		Age:     age,
		Address: userData[2],
	} //your code here!

}

func main() {
	user1 := ConvertData("Edo,20,Jakarta")
	fmt.Printf("%+v\n", user1)

	user2 := ConvertData("Budi,30,Bandung")
	fmt.Printf("%+v\n", user2)
}
