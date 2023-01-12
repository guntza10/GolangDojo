package main

import (
	"fmt"
	"strings"
)

/*
function => encapsulate code ที่ logic ทำ action บางอย่างเดียวกัน
- function จะถูก execute ก็ต่อเมื่อมีการ call
- function สามารถ call ได้หลายครั้ง
- function ช่วยลด duplicated code

func name(valueName type) {

}
*/

/*
package level variables
- กำหนด variable ไว้ด้านนอกสุดของทุก function
- ทุก function สามารถ access ได้หมด
- ทุกอย่างใน file ที่อยู่ใน package เดียวกัน ก็สามารถ access ได้หมด
- ไม่สามารถใช้ Syntactic Sugar short hand ได้

local variables
- กำหนดใน function หรือ block state
- สามารถ access ได้แค่ใน function หรือ block state นั้นๆที่ declare เท่านั้น
*/
const conferenceTickets int = 50

var remainingTickets uint = 50
var conferenceName = "Go Conference"
var bookings = []string{}

func main() {

	greetUsers()

	for {

		// get data from function return multiple value
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			firstNames := printFirstNames()
			fmt.Printf("The first names %v\n", firstNames)

			if remainingTickets == 0 {
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

/*
function return
- กำหนด type ของ function ให้ตรงกับ type ของ value ที่ต้องการ return

	funct name() type {
		return v
	}
*/
func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

/*
function return multiple value
- กำหนด type ของ function ให้ตรงกับ type ของ value ที่ต้องการ return

	funct name() (type,type,type) {
		return v1,v2,v3
	}
*/
func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}