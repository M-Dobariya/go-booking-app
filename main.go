package main

import (
	"fmt"
	"strconv"
	"strings"
)

type UserInfo struct {
	firstName           string
	lastName            string
	userEmail           string
	ticketsWantedByUser int
}

const totalTickets int = 100

var remainingTickets int = totalTickets
var ticketBookerUserDetails []UserInfo
var tickerBookerFirstNames []string

// var userInfo UserInfo = {
// 	var firstName string
// 	var lastName string
// 	var userEmail string
// 	var ticketsWantedByUser uint
// }

func main() {
	greetUser()
	for remainingTickets > 0 {
		firstName, lastName, userEmail, ticketsWantedByUser := getUserInfo()

		remainingTickets = remainingTickets - ticketsWantedByUser

		userInfo := UserInfo{
			firstName:           firstName,
			lastName:            lastName,
			userEmail:           userEmail,
			ticketsWantedByUser: ticketsWantedByUser,
		}

		ticketBookerUserDetails = append(ticketBookerUserDetails, userInfo)
		tickerBookerFirstNames = append(tickerBookerFirstNames, firstName)
	}

	fmt.Println(ticketBookerUserDetails)
}

func getUserInfo() (string, string, string, int) {
	var firstName string
	var lastName string
	var userEmail string
	var ticketsWantedByUser int

	var isFirstNameCorrect bool
	var isLastNameCorrect bool
	var isUserEmailCorrect bool
	var isTicketQuantityCorrect bool

FIRST_NAME:
	fmt.Println("Enter The First Name Of The User")
	firstName = takeUserInput()

	isFirstNameCorrect = validateUserInput(firstName, "FIRST_NAME")

	if !isFirstNameCorrect {
		showErrorLog("FIRST_NAME")
		goto FIRST_NAME
	}
LAST_NAME:
	fmt.Println("Enter The Last Name Of The User")
	lastName = takeUserInput()

	isLastNameCorrect = validateUserInput(lastName, "LAST_NAME")

	if !isLastNameCorrect {
		showErrorLog("LAST_NAME")
		goto LAST_NAME
	}
EMAIL:
	fmt.Println("Enter The Email Of The User")
	userEmail = takeUserInput()

	isUserEmailCorrect = validateUserInput(userEmail, "EMAIL")

	if !isUserEmailCorrect {
		showErrorLog("EMAIL")
		goto EMAIL
	}
TICKETS_WANTED:
	fmt.Println("Enter Number Of Ticket Wanted")
	ticketsWantedByUser, _ = strconv.Atoi(takeUserInput())

	isTicketQuantityCorrect = validateUserInput(ticketsWantedByUser, "TICKETS_WANTED")

	if !isTicketQuantityCorrect {
		showErrorLog("TICKETS_WANTED")
		goto TICKETS_WANTED
	}

	return firstName, lastName, userEmail, ticketsWantedByUser
}

func showErrorLog(errorEvent string) {
	fmt.Printf("Error in %s value, once more\n", errorEvent)
}

func greetUser() {
	fmt.Printf("*************************** Welcome Bolte Kya Bidu.... ***************************\n\n")
}

func validateUserInput(inputValue interface{}, fieldToValidate string) bool {

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

func takeUserInput() string {
	var valueRead string

	fmt.Print("Enter The Value: ")
	fmt.Scan(&valueRead)

	return string(valueRead)
}
