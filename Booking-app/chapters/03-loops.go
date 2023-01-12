package main

// import multiple packages
import (
	"fmt"
	"strings"
)

func loops() {

	/*
		loop ใน Go มีแค่ for loop แค่อันเดียวเท่านั้น

		1.Infinity loop

		- สามารถใส่ condition ให้ loop เพื่อเช็คว่าจะ loop ต่อหรือไม่ (ไม่ใส่ condition ก็ได้ มันก็จะ infinity loop)
		- ถ้า condition = true(loop ต่อ),ถ้า condition = false(loop หยุด)

		for (condition) {

		}

		2. For-Each loop

		for 'index', 'element' := range 'array' {

		}

		Note:
		- หยุด loop ด้วย break (terminate the for loop)
		- ให้ทำ loop ต่อไปด้วย continue (มันจะ skip loop ต่อไปเลยโดยที่ไม่สน statement ที่เหลือ)
	*/
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// asking for user input
		fmt.Println("Enter Your First Name: ")
		fmt.Scanln(&firstName)

		fmt.Println("Enter Your Last Name: ")
		fmt.Scanln(&lastName)

		fmt.Println("Enter Your Email: ")
		fmt.Scanln(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scanln(&userTickets)

		// book ticket in system
		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		// print only first names
		firstNames := []string{}
		for _, booking := range bookings { // _ => blank identifier => ignore variable that don't want to use
			var names = strings.Fields(booking) // splits the string with the white space
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names %v\n", firstNames)
	}
}
