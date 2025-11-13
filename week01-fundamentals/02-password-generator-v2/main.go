package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func printBanner() {
	fmt.Println("Password generator v2")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println()
}

func getIntInput(prompt string, min int) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input please try again.")
			continue
		}

		//trim whitespace and newline
		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid number.")
			continue
		}

		if value < min {
			fmt.Printf("Value must be at least %d. Please try again.\n", min)
			continue
		}

		return value
	}
}

func getStringInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input. Please try again.")
			continue
		}

		input = strings.TrimSpace(input)

		if input == "" {
			fmt.Println("Input cannot be empty. Please try again")
			continue
		}

		return input
	}
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
		return "Very Strong"
	} else if score >= 4 {
		return "Strong"
	} else if score >= 3 {
		return "Medium"
	} else {
		return "Weak"
	}
}

func buildCharset() string {
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"
	symbols := "!@#$%^&*()-_=+[]{}|;:,.<>?"

	return lowercase + uppercase + numbers + symbols
}

func savePasswordToFile(entry string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		homeDir = "."
	}

	saveDir := filepath.Join(homeDir, "Passwords")
	filePath := filepath.Join(saveDir, "passwords.txt")

	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		if err := os.MkdirAll(saveDir, 0755); err != nil {
			return fmt.Errorf("Failed to create directory: %v", err)
		}
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Failed to open file: %v", err)
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	entryWithTime := fmt.Sprintf("[%s] %s\n", timestamp, entry)

	if _, err := file.WriteString(entryWithTime); err != nil {
		return fmt.Errorf("Failed to write to file: %v", err)
	}
	return nil
}

func main() {
	rand.Seed(time.Now().UnixNano())

	printBanner()

	length := getIntInput("Password length (min 8): ", 8)
	fmt.Printf("length set to: %d\n\n", length)

	quantity := getIntInput("How many passwords (min 1): ", 1)
	fmt.Printf("Will generate %d password(s)\n", quantity)

	charset := buildCharset()

	fmt.Println("\n" + strings.Repeat("-", 50))

	for i := 1; i <= quantity; i++ {
		label := getStringInput(fmt.Sprintf("\n Enter label for password #%d (e.g. github, gmail): ", i))

		password := generatePassword(length, charset)
		strength := checkPasswordStrength(password)

		fmt.Printf("Generated for '%s':%s\n", label, password)
		fmt.Printf("Length: %d | Strength: %s\n", len(password), strength)

		entry := fmt.Sprintf("%s - %s", label, password)
		if err := savePasswordToFile(entry); err != nil {
			fmt.Println("Error saving password:", err)
		} else {
			homeDir, _ := os.UserHomeDir()
			savePath := filepath.Join(homeDir, "Passwords", "passwords.txt")
			fmt.Printf("Saved to: %s\n", savePath)
		}
	}

	fmt.Println(strings.Repeat("-", 50))

	fmt.Println("\n Password generation completed!")
}
