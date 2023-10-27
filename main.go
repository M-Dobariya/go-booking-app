package main

import (
	"fmt"
	"strconv"
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
	GreetUser()
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
	firstName = TakeUserInput()

	isFirstNameCorrect = ValidateUserInput(firstName, "FIRST_NAME")

	if !isFirstNameCorrect {
		ShowErrorLog("FIRST_NAME")
		goto FIRST_NAME
	}
LAST_NAME:
	fmt.Println("Enter The Last Name Of The User")
	lastName = TakeUserInput()

	isLastNameCorrect = ValidateUserInput(lastName, "LAST_NAME")

	if !isLastNameCorrect {
		ShowErrorLog("LAST_NAME")
		goto LAST_NAME
	}
EMAIL:
	fmt.Println("Enter The Email Of The User")
	userEmail = TakeUserInput()

	isUserEmailCorrect = ValidateUserInput(userEmail, "EMAIL")

	if !isUserEmailCorrect {
		ShowErrorLog("EMAIL")
		goto EMAIL
	}
TICKETS_WANTED:
	fmt.Println("Enter Number Of Ticket Wanted")
	ticketsWantedByUser, _ = strconv.Atoi(TakeUserInput())

	isTicketQuantityCorrect = ValidateUserInput(ticketsWantedByUser, "TICKETS_WANTED")

	if !isTicketQuantityCorrect {
		ShowErrorLog("TICKETS_WANTED")
		goto TICKETS_WANTED
	}

	return firstName, lastName, userEmail, ticketsWantedByUser
}
