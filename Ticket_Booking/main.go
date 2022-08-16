//============ Using the newly created package shared.go =============//

package main

import (
	"Ticket_Booking/shared"
	"fmt"
	"strings"
)

const conferenceTickets int = 50 // Package level variables (Declared globally)
var remainingTickets int = 50
var conferenceName = "Go Conference"
var bookings = []string{}

func main() {

	greetuser()

	for {
		firstName, lastName, email, userTickets := userInput()
		// validate user input
		isValidName, isValidEmail, isValidTicketNumber := shared.ValidateUserDetails(firstName, lastName, email, userTickets, remainingTickets) //Exporting a member function from the additional imported package.

		if isValidName && isValidEmail && isValidTicketNumber {
			// book ticket in system
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			// print only first names
			firstNames := getFirstname()
			fmt.Printf("The first names %v\n", firstNames)

			// exit application if no tickets are left
			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}

}

func greetuser() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are available.\n", conferenceName, conferenceTickets, remainingTickets)
}

func getFirstname() []string {
	userfirstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking) //To seperate the each field of the string
		userfirstNames = append(userfirstNames, names[0])
	}
	return userfirstNames

}

func userInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}
