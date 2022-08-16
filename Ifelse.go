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
		var firstName string
		var lastName string

		fmt.Printf("Welcome to the %v ticket booking platform\n", conferenceName)
		fmt.Println("Get your tickets to attend")
		fmt.Println("Enter the first and last name.")
		fmt.Scan(&firstName)
		fmt.Scan(&lastName)

		fmt.Println("Enter the number of tickets for booking.")
		var userTickets int
		fmt.Scan(&userTickets)
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("We have total %v tickets and %v are available for booking\n", conferenceTickets, remainingTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We have only %v tickets available, You can't book %v\n", remainingTickets, userTickets)
			break
		}

		var bookings []string
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you for the ticket booking for %v!, Available tickets for the bookings are %v.\n", conferenceName, remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking) //To seperate the each field of the string
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of the bookings are: %v\n", firstNames)
		if remainingTickets == 0 {
			fmt.Println("You can't book more tickets")
			break
		}
	}
}
