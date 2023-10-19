package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"unicode"
)

func main() {

	var password string
	var special_characters = []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')'}
	var errors = []string{}

	// For loop
	for {
		clear_output()

		if len(errors) > 0 {
			fmt.Println("\n\t\t\tError list:")
			for _, error := range errors {
				fmt.Print(error)
			}
			errors = []string{}
			fmt.Print("\n\n")
		}

		fmt.Println("Passwords length must be at least 8 symbols long\nand must contain:\n1)Uppercase\n2)Lowercase\n3)Number\n4)Special symbol")
		fmt.Print("Please, enter your password: ")
		_, err := fmt.Scan(&password)

		uppercase := false
		lowercase := false
		number := false
		special_char := false

		if err == nil {
			if len(password) >= 8 {
				for _, value := range password {
					switch {
					case IsInside(special_characters, value):
						special_char = true
						continue
					case unicode.IsUpper(value):
						uppercase = true
						continue
					case unicode.IsLower(value):
						lowercase = true
						continue
					case unicode.IsNumber(value):
						number = true
						continue
					}
				}
				if uppercase && lowercase && number && special_char {
					fmt.Println("Success")
					break
				} else {
					errors = append(errors, "Uppercase: "+strconv.FormatBool(uppercase)+"\nLowercase: "+strconv.FormatBool(lowercase)+"\nNumber: "+strconv.FormatBool(number)+"\nSpecial character: "+strconv.FormatBool(special_char)+"\n")
				}
			} else {
				errors = append(errors, "Password must be at least 8 symbols long")
			}
		} else {
			fmt.Println(err)
		}
	}

}

// Checks if rune in the list
func IsInside(list []rune, char rune) bool {
	for _, value := range list {
		if value == char {
			return true
		}
	}
	return false
}

// Executes "clear" command
func clear_output() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
