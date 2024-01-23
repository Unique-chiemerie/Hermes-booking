package main

import "fmt"

func main() {
	var conferenceName = "GO conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("confrence is %T, remaining tickets is %T, conferencename is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have a total of %v, tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Book tickets now")

	var userName string
	var userTickets int
	//asking the user their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("user %v booked %v tickets\n", userName, userTickets)

}
