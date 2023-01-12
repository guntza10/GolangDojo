package chapters

import "fmt"

func array() {
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	conferenceName := "Go Conference"
	/*
		array => ต้องกำหนด size ให้มัน พร้อมกับ datatype [size]datatype
		- การ assign value => var bookings = [50]string ,var bookings = [50]string {"gunt","gunt2","gunt3"}
		- การ declare array => var bookings [50]string
		- การ assign value ให้ element ใน array => bookings[0] = "gunt", bookings[10] = "gunt2" (ต้อง ref index ของ array)
		- การอ้างอิงถึง element => bookings[0],bookings[10] (ref ด้วย index ของ array)
		- length array => len(array)
		- ถ้า index เกิน size array => มันจะ out of range error
	*/

	/*
		slice => ก็คือ array แบบที่ไม่ต้องกำหนด size จะมีความ dynamic มากกว่า
		- การ assign value => var bookings = []string ,var bookings = []string {"gunt","gunt2","gunt3"}
		- การ declare slice => var bookings []string
		- การ assign value ให้ element ใน slice โดยไม่ได้ ref index => append(slice, value) (จะเพิ่ม element เข้าไปที่ index สุดท้ายของ slice)
		- การอ้างอิงถึง element => bookings[0],bookings[10] (ref ด้วย index ของ slice)
		- length slice => len(slice)
	*/
	bookings := []string{}

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)

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
	fmt.Printf("These are all our bookings: %v\n", bookings)
}
