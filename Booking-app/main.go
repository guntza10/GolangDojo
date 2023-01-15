package main

/*
	package
	- Go program ถูกจัดการใน package
	- package คือ collection ของ Go file
	- ทุกอย่างใน file ที่อยู่ใน package เดียวกัน ก็สามารถ access ได้หมด
*/

/*
	import another package in main
	- path import <module name>/<folder package name>
	- วิธีเรียกเอา <folder package name> มาดอทใช้ function,value
	- ต้อง export value,function จาก another package ด้วย ถ้าต้องการจะเอามาใช้ที่นี่ (เปลี่ยนให้เป็น capital case)
*/
import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

/*
WaitGroup => มี 3 function
- Add(no) => จะเป็นการบอก main thread ว่ามีจำนวน thread เท่าไรที่มันต้องรอและต้อง execute ก่อนที่จะ สร้าง new thread
- Wait() => จำเป็นต้อง execute ไว้ตอนสุดท้ายของ main thread มันจะรอจนกว่าทุก thread จะ done (WaitGroup list เป็น 0 ไม่เหลือ thread)
จะทำงานก่อนที่ app จะ exit terminate ทุกอย่าง
- Done() => จำเป็นต้อง execute ไว้ตอนสุดท้ายของ separated thread(Goroutine thread) เมื่อถูก call มันจะทำการลดจำนวน WaitGroup ลง 1
(remove thread ที่อยู่ใน WaitGroup list ออกไป 1)
*/
// use Synchronizing Goroutine
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	// for {
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {

		bookTicket(userTickets, firstName, lastName, email)

		// บอก main thread ว่ามี 1 thread ที่ต้องรอและ execute
		wg.Add(1)
		// use Goroutine => separate thread execute
		go sendTicket(userTickets, firstName, lastName, email)

		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			// break
		}
	} else {
		if !isValidName {
			fmt.Println("first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
	}
	//}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// simulate if process take some time
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}

/*
Goroutine - Concurrency
- ปกติเวลาเรา execute app มันจะรันเป็น single thread ตั้งแต่ต้นจนจบตามลำดับ
- ถ้ามี process ใดที่ใช้เวลาในการ execute นาน จะทำให้ thread ติด block
- ทำให้ code process ในบรรทัดต่อไปต้องรอ process ก่อนหน้าให้ execute เสร็จก่อน
- เราจะใช้ Goroutine(keyword go) เข้ามาช่วย
- โดยที่ Goroutine จะแยก thread ของ process ที่ใช้เวลานานออกมา execute ต่างหาก (run in background)
*/

/*
Synchronizing the Goroutines
- ปัญหาคือ main thread execute เสร็จก่อนที่ตัว Goroutine จะเริ่มและ execute code
- main thread ไม่ได้รอ Goroutine execute
- พอ main thread execute เสร็จมัน app ก็จบการทำงานพร้อมกับ terminate ทุก thread ทิ้ง
- เราใช้ `WaitGroup` เข้ามาช่วย เพื่อบอกให้ตัว main thread รอ Goroutine execute
*/
