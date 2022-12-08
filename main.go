package main

import (
	"fmt"
	"sync"
)

const showTickets int = 100

var showName = "Go Show"
var remainingTickets uint = 100
var bookings = make([]UserData, 0)

type UserData struct {
	firstName     string
	lastName      string
	email         string
	ticketsNumber uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	//for {
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(firstName, lastName, email, userTickets)

		wg.Add(1)
		go sendTickets(userTickets, email)

		if remainingTickets == 0 {
			fmt.Printf("For now there are no more available booking for the %v.\n", showName)
			//break
		}

	} else {
		if !isValidName {
			fmt.Println("First or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Not a valid email adress")
		}
		if !isValidTicketNumber {
			fmt.Printf("We have %v tickets remaining. You can't book %v tickets.\n", remainingTickets, userTickets)
		}

	}
	//}
	wg.Wait()
}
