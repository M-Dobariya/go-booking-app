package main

import (
	"fmt"
	"strings"
)

func ShowErrorLog(errorEvent string) {
	fmt.Printf("Error in %s value, once more\n", errorEvent)
}

func GreetUser() {
	fmt.Printf("*************************** Welcome Bolte Kya Bidu.... ***************************\n\n")
}

func ValidateUserInput(inputValue interface{}, fieldToValidate string) bool {

	valueOfStrVariable, isString := inputValue.(string)
	valueOfIntVariable, isInt := inputValue.(int)

	if isString {
		switch fieldToValidate {
		case "FIRST_NAME", "LAST_NAME":
			return len(valueOfStrVariable) >= 1
		case "EMAIL":
			return strings.Contains(valueOfStrVariable, "@")
		default:
			return false
		}
	} else if fieldToValidate == "TICKETS_WANTED" && isInt {
		return valueOfIntVariable > 0 && valueOfIntVariable <= remainingTickets
	}

	return false
}

func TakeUserInput() string {
	var valueRead string

	fmt.Print("Enter The Value: ")
	fmt.Scan(&valueRead)

	return string(valueRead)
}
