package main

import (
	"fmt"
	"strings"
)

const showTickets int = 100

var showName = "Go Show"
var remainingTickets uint = 100
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(firstName, lastName, email, userTickets)
			printFirstNames()

			if remainingTickets == 0 {
				fmt.Printf("For now there are no more available booking for the %v.\n", showName)
				break
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

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", showName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", showTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func printFirstNames() {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("The first names of booking are: %v\n", firstNames)
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, showName)
}
