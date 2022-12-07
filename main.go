package main

import "fmt"

func main() {
	var showName = "Go Show"
	const showTickes = 100
	var remaingTickets uint = 100

	fmt.Printf("Welcome to %v booking application!\n", showName)
	fmt.Println("We have a total of", showTickes, "tickes and", remaingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend!")

	var fullName string
	var email string
	var userTickets uint

	fmt.Println("Enter your fisrt name: ")
	fmt.Scan(&fullName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets to book: ")
	fmt.Scan(&userTickets)

	remaingTickets = remaingTickets - userTickets

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", fullName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remaingTickets, showName)
}
