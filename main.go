package main

import (
	"fmt" //imports format package, used for things like Print, Println, etc
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	//fmt.Println("Welcome to our", conferenceName, "booking application") //using normal println
	fmt.Printf("Welcome to our %v booking application\n", conferenceName) //using Printf with a placehoulder %v
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
