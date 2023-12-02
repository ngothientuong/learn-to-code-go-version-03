package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs := []byte(s)
	new_bs, err := bcrypt.GenerateFromPassword(bs, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(new_bs)
}
