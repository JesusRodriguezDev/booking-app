package main

import (
	"fmt" //imports format package, used for things like Print, Println, etc
)

func main() {
	conferenceName := "Go Conference" //simplified alternative variable declaration, cannot be used with const, only with var or if you want to define the data type
	const conferenceTickets = 50      //in Go variables have to either be assigned a value when created to imply data type, or declare the datatype i.e. var userName string
	var remainingTickets uint = 50    //uint data type for possitive numbers.  This prevents negative numbers to be used

	//fmt.Println("Welcome to our", conferenceName, "booking application") //using normal println
	fmt.Printf("Welcome to our %v booking application\n", conferenceName) //using Printf with a placeholder %v
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //using & at the beginning of the variable name to point to the memory address where the variable is stored. The program will then wait for input and use Scan to get that value of that input and assign the it to the variable

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets to purchase")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

}
