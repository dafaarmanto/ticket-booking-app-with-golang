package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	firstName, lastName, email, userTickets := getUserInputs()
	validatedName, validatedEmail, validatedTickets := helper.ValidateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

	// Validation
	if validatedName && validatedEmail && validatedTickets {
		// Logic
		bookTicket(userTickets, firstName, lastName, email)

		// Send Tickets -- WaitGroup
		wg.Add(1)
		go sendTickets(firstName, lastName, email, userTickets)

		firstNames := getFirstNames()
		fmt.Printf("The first names of the attendees are: %v\n", firstNames)
		fmt.Print("\nPlease wait while we send your tickets...\n")

		if remainingTickets == 0 {
			fmt.Println("SOLD OUT! Come back next year!")
		}
	} else {
		// Error Handling
		if !validatedName {
			fmt.Println("Please enter a valid name")
		}

		if !validatedEmail {
			fmt.Println("Please enter a valid email")
		}

		if !validatedTickets {
			fmt.Printf("Sorry, we only have %v tickets remaining\n", remainingTickets)
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("\nWelcome to the %v\n", conferenceName)
	fmt.Printf("There are %v tickets available and %v remaining\n", conferenceTickets, remainingTickets)
	fmt.Print("Get yout tickets here\n\n")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// User Input
	fmt.Print("Enter your name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("How many tickets would you like to purchase: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	User := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, User)
	fmt.Printf("\nList of attendees: %v\n", bookings)

	fmt.Printf("%v %v, your tickets are booked for %v.\nYou will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v tickets remaining\n\n", remainingTickets)
}

func sendTickets(firstName string, lastName string, email string, userTickets uint) {
	time.Sleep(time.Second * 10)
	var ticket = fmt.Sprintf("Dear %v %v,\n\n You have booked %v tickets for the %v.\n\n Regards,\n The Organisers", firstName, lastName, userTickets, conferenceName)
	fmt.Print("#############################################################\n")
	fmt.Printf("Sending ticket:\n\n %v \n\nTo E-mail: %v\n", ticket, email)
	fmt.Print("\nThanks for purchasing tickets!\nEnjoy the conference!\n")
	fmt.Print("#############################################################\n")
	wg.Done()
}
