package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of ticket: ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v  \n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remainning for %v\n", remainingTickets, conferenceName)

}