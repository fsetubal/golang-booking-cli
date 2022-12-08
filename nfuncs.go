package main

import (
	"fmt"
	"strings"
	"time"
)

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", showName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", showTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
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

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTicket(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:     firstName,
		lastName:      lastName,
		email:         email,
		ticketsNumber: userTickets,
	}

	bookings = append(bookings, userData)
	// fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, showName)
}

func sendTickets(userTickets uint, email string) {
	time.Sleep(10 * time.Second)
	fmt.Println("################################")
	fmt.Printf("%v ticket(s) sent to %v\n", userTickets, email)
	wg.Done()
}
