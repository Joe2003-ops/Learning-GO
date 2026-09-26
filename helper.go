package main

import "strings"
func validateUserInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool){
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= RemainingTickets
	
	return isValidName, isValidEmail, isValidTicketNumber
}