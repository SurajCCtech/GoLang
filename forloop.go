package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		const conferenceTickets = 100
		var conferenceName = "GO Conference"
		var remainingTickets = 100

		fmt.Printf("Welcome to the %v ticket booking platform\n", conferenceName)
		fmt.Printf("We have total %v tickets and %v are available for booking\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets to attend")
		fmt.Println("Enter the Username and the number of tickets to book")

		var firstName string
		var lastName string
		var userTickets int

		fmt.Scan(&firstName)
		fmt.Scan(&lastName)
		fmt.Scan(&userTickets)

		// var bookings [100]string      //Array of fixed size
		// bookings[0] = firstName + " " + lastName

		var bookings []string //Slice also written as ""var bookings = []string{}""
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("whole slice is: %v\n", bookings)
		fmt.Printf("First element of slice: %v\n", bookings[0])
		fmt.Printf("Type of slice: %T\n", bookings)

		fmt.Printf("User %v %v has booked %v tickets\n", firstName, lastName, userTickets)
		remainingTickets = remainingTickets - userTickets
		fmt.Printf("Thank you for the ticket booking for %v!, Available tickets for the bookings are %v.\n", conferenceName, remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking) //To seperate the each field of the string
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of the bokkings are: %v\n", firstNames)
	}
}
