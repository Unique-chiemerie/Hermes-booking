package main

import "fmt"

func main() {
	var conferenceName = "GO conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have a total of %v, tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Book tickets now")
	var userName = "something"
	//asking the user their name

	fmt.Println(userName)

}
