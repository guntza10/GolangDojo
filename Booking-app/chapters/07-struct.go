package chapters

import (
	"fmt"
	"strings"
)

/*
struct
- ก็คือ class ในภาษาอื่นๆ (doesn't support inheritance)
- define key pair กับ value ได้หลาย type (property)
- declare struct นอก function
*/
const conferenceTickets3 int = 50

var remainingTickets3 uint = 50
var conferenceName3 = "Go Conference"
var bookings3 = make([]User, 0)

// declare struct
type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func structs() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets3 == 0 {
				// end program
				break
			}
		} else {
			if !isValidName {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}
}

func printFirstNames3() []string {
	firstNames := []string{}

	for _, booking := range bookings3 {
		// get property from struct ใช้ dot
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput3() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func validateUserInput3(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets3
	return isValidName, isValidEmail, isValidTicketNumber
}

func greetUsers3() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName3, conferenceTickets3, remainingTickets3)
}

func bookTicket3(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets3 = remainingTickets3 - userTickets

	// define struct value
	var user = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings3 = append(bookings3, user)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets3, conferenceName3)
}
