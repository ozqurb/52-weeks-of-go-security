package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Banner
	fmt.Println("Password Generator v1.0")
	fmt.Println("=========================")
	fmt.Println()

	rand.Seed(time.Now().UnixNano())

	var length int
	var quantity int

	fmt.Print("specify password length:")
	fmt.Scan(&length)

	fmt.Print("How many passwords will be generated:")
	fmt.Scan(&quantity)

	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>?"

	allChars := lowercase + uppercase + numbers + symbols

	for q := 1; q <= quantity; q++ {
		password := ""
		for i := 0; i < length; i++ {
			randomIndex := rand.Intn(len(allChars))
			password += string(allChars[randomIndex])
		}
		fmt.Printf("Generated Password: %s\n", password)
		fmt.Printf("Length: %d characters\n", len(password))
	}

	fmt.Println()
	fmt.Println("Password generated")
}
