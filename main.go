package main

import (
	"fmt"     //imports format package, used for things like Print, Println, etc
	"strings" //imports strings package, used for string manipulation
)

func main() {
	conferenceName := "Go Conference" //simplified alternative variable declaration, cannot be used with const, only with var, or if you want to define the data type
	const conferenceTickets = 50      //in Go variables have to either be assigned a value when created to imply data type, or declare the datatype i.e. var userName string
	var remainingTickets uint = 50    //uint data type for possitive numbers.  This prevents negative numbers to be used
	//var bookings []string //this creates a slice, which is more effecient than an array as it allows for dynamic size so a size does not have to be declared
	bookings := []string{} //alternative way of declaring slice variable

	greetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) //using & at the beginning of the variable name to point to the memory address where the variable is stored. The program will then wait for input and use Scan to get that value of that input and assign the it to the variable

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets to purchase")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings { //the underscore used after for, instead of index, declares an anonymous variable, since we just need it to start up the for loop, but won't use it later
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of the bookings are:  %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year!")
				break //break ends the loop
			}
		} else {
			if !isValidName {
				fmt.Println("first name of last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets entered is invalid")
			}
		}
	}
}

func greetUsers(confName string, confTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to our %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}
