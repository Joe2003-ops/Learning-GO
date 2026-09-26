package main

import (
	"fmt"
	"sync"
	"time"
)

const ConferenceTickets = 50
var ConferenceName = "Go Conference"
var RemainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main(){
	//Landing terminal
	greetUsers()
	
	// Loop to book the tickets

	// Call the user input function
	firstName, lastName, email, userTicket := getUserInput()

	// Call function validUserInput
	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket)

	if isValidName && isValidEmail && isValidTicketNumber {
		// Call book ticket function
		bookTicket(userTicket, firstName, lastName, email)
		
		// Adding a wait thread before program terminates
		wg.Add(1)
		// Calling the send ticket function
		go sendTicket(userTicket, firstName, lastName, email)

		// Call function print first name
		firstNames := getFirstName()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		var noTicketsRemaining bool = RemainingTickets == 0
		if noTicketsRemaining{
			// end the program
			fmt.Println("Our conference is booked out. Come back next year.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("First name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is invalid")
		}
	}
	wg.Wait()
}	

func greetUsers(){
	//Landing terminal
	fmt.Printf("Welcome to %v booking application\n", ConferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", ConferenceTickets, RemainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTicket
}

func bookTicket (userTicket uint, firstName string, lastName string, email string) {
	RemainingTickets = RemainingTickets - userTicket

	// Create a map for a user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booked %v tickets. You'll recieve a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v.\n", RemainingTickets, ConferenceName)
}

func sendTicket(userTicket uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##########")
	wg.Done()
}