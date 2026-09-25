package main

import (
	"fmt"
	"strings"
)

func main(){
	ConferenceName := "Go Conference"
	const ConferenceTickets = 50
	var RemainingTickets uint = 50
	var bookings []string

	//Landing terminal
	fmt.Printf("Welcome to %v booking application\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", ConferenceTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend")
	
	// Loop to book the tickets
	for RemainingTickets > 0 && len(bookings) < 50{
		var firstName string
		var lastName string
		var email string
		var userTicket uint

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

		if userTicket < RemainingTickets {
			RemainingTickets = RemainingTickets - userTicket
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %v %v for booked %v tickets. You'll recieve a confirmation email at %v\n", firstName, lastName, userTicket, email)
			fmt.Printf("%v tickets remaining for %v.\n", RemainingTickets, ConferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			var noTicketsRemaining bool = RemainingTickets == 0
			if noTicketsRemaining{
				// end the program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			fmt.Printf("We only have  %v tickets remaining, so you can't book %v tickets\n", RemainingTickets, userTicket)
		}
	}
}	