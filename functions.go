//============ Functions / Functions with Parameters / Multiple return values / Package level variables =============//

package main

import (
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
		isValidName, isValidEmail, isValidTicketNumber := validateUserDetails(firstName, lastName, email, userTickets)

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

func validateUserDetails(firstName, lastName, email string, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
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
