package main

import "fmt"

func main(){
	ConferenceName := "Go Conference"
	const ConferenceTickets = 50
	var RemainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", ConferenceTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTicket uint
	var bookings [50]string

	// Ask user for their first name
	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	// Ask user for their first name
	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	// Ask user for their Email address
	fmt.Println("Enter your email address")
	fmt.Scan(&email)

	// Ask user the numbeer of tickets
	fmt.Println("Enter number of tickets")
	fmt.Scan(&userTicket)

	RemainingTickets = RemainingTickets-userTicket
	
	fmt.Printf("Thank you %v %v for booked %v tickets. You'll recieve a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v.\n", RemainingTickets, ConferenceName)

	bookings[0] = firstName + " " + lastName
	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("Array Type: %T\n", bookings)
	fmt.Printf("Array Length: %v\n", len(bookings))
}