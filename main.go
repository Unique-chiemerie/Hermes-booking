package main

import "fmt"

func main() {
	conferenceName := "GO conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have a total of %v, tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Book tickets now")

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//asking the user their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)
	//asking their last name
	fmt.Println("Enter your Last name:")
	fmt.Scan(&lastName)
	//asking for thier email
	fmt.Println("Enter your Email address:")
	fmt.Scan(&email)
	//number of tickets
	fmt.Println("Enter your number of Tickets to get:")
	fmt.Scan(&userTickets)
	remainingTickets = remainingTickets - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
