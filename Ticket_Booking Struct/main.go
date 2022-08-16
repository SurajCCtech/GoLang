//============ Using the map and the strconv package =============//

package main

import (
	"Ticket_Booking/shared"
	"fmt"
	"time"
)

const conferenceTickets int = 50

var remainingTickets int = 50
var conferenceName = "Go Conference"
var bookings = make([]UserData, 0)

type UserData struct { //structure declaration
	firstName       string
	lastName        string
	email           string
	numberOfTickets int
}

func main() {

	greetuser()

	for {
		firstName, lastName, email, userTickets := userInput()
		// validate user input
		isValidName, isValidEmail, isValidTicketNumber := shared.ValidateUserDetails(firstName, lastName, email, userTickets, remainingTickets) //Exporting a member function from the additional imported package.

		if isValidName && isValidEmail && isValidTicketNumber {
			// book ticket in system
			bookTickets(userTickets, firstName, lastName, email)

			// Started new thread for printing or sending the ticket
			go sendTicket(userTickets, firstName, lastName, email)

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
		userfirstNames = append(userfirstNames, booking.firstName) // Accessing the structure members with the dot operators
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

func bookTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{ // Structure definition using the object
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of the bookings %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets int, firstName, lastName, email string) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets booked by %v %v\n", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending ticket: \n%v \nto email address%v\n", ticket, email)
	fmt.Println("#######################")
}
