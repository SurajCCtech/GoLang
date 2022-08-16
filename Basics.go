package main

import "fmt"

func main() {
	const conferenceTickets = 100
	var conferenceName = "GO Conference"
	var remainingTickets = 100

	fmt.Printf("Welcome to the %v ticket booking platform\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are available for booking\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")
	fmt.Println("Enter the Username and the number of tickets to book")

	var userName string
	var userTickets int

	fmt.Scan(&userName)
	fmt.Scan(&userTickets)

	fmt.Printf("User %v has booked %v tickets\n", userName, userTickets)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you for the ticket booking for %v!, Available tickets for the bookings are %v.", conferenceName, remainingTickets)
}
