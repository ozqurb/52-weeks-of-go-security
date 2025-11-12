package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func printBanner() {
	fmt.Println("Password Generator v2.0")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
}

func getIntInput(prompt string) (int, error) {
	var value int
	fmt.Print(prompt)

	_, err := fmt.Scan(&value)
	if err != nil {
		return 0, fmt.Errorf("invalid input: please enter a number")
	}

	if value <= 0 {
		return 0, fmt.Errorf("value must be greater than 0")
	}

	return value, nil
}

func generatePassword(length int, charset string) string {
	password := ""
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		password += string(charset[randomIndex])
	}
	return password
}

func checkPasswordStrength(password string) string {
	hasLower := strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz")
	hasUpper := strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hasNumber := strings.ContainsAny(password, "0123456789")
	hasSymbol := strings.ContainsAny(password, "!@#$%^&*()-_=+[]{}|;:,.<>?")

	score := 0

	if hasLower {
		score++
	}
	if hasUpper {
		score++
	}
	if hasNumber {
		score++
	}
	if hasSymbol {
		score++
	}

	length := len(password)

	if length >= 16 {
		score++
	}

	if score >= 5 {
		return "very strong"
	} else if score >= 4 {
		return "strong"
	} else if score >= 3 {
		return "medium"
	} else {
		return "weak"
	}
}

func buildCharset() string {
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>?"

	return lowercase + uppercase + numbers + symbols
}

func main() {
	rand.Seed(time.Now().UnixNano())

	printBanner()

	length, err := getIntInput("Password length (min 8): ")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if length < 8 {
		fmt.Println("Warning: Password length should be at least 8 characters")
	}

	quantity, err := getIntInput("How many passwords")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	charset := buildCharset()

	fmt.Println("\n" + strings.Repeat("-", 50))
	for i := 1; i <= quantity; i++ {
		password := generatePassword(length, charset)
		strength := checkPasswordStrength(password)

		fmt.Printf("\n#%d Password: %s\n", i, password)
		fmt.Printf("   Length: %d | Strength: %s\n", len(password), strength)
	}

	fmt.Println(strings.Repeat("-", 50))

	fmt.Println("\n Password generation completed!")
}
