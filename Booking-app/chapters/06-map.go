package chapters

import (
	"fmt"
	"strconv"
	"strings"
)

const conferenceTickets int = 50

/*
map => data type ที่เป็น object เหมือนภาษาอื่นๆ ที่จะเก็บข้อมูลเป็นก้อนด้วย key(property) map กับ value
- map[type_key]type_value (เป็นแค่การ declare type)
- ทุก key จะต้องมี data type เหมือนกัน
- ทุก value จะต้องมี data type เหมือนกัน
- create empty map => make(map[type_key]type_value)
*/
var remainingTickets uint = 50
var conferenceName = "Go Conference"

/*
วิธีการสร้าง slice ของ map (list of map)
- จะไม่เหมือนทั่วไป ([]type => []map[type_key]type_value แบบนี้ผิด)
- ต้องสร้างผ่าน make => make([]map[type_key]type_value, initail_size)
*/
var bookings = make([]map[string]string, 0)

func main() {

	greetUsers()

	for {

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

func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}
	return firstNames
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

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create user map
	var user = make(map[string]string)
	/*
		define value ให้ map
		- mapData[key] = value
	*/
	user["firstName"] = firstName
	user["lastName"] = lastName
	user["email"] = email
	// strconv เอาไว้ convert type ต่างๆให้เป็น string ดูต่อได้ที่ document
	user["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
