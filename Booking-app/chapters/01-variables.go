package main

import "fmt"

func main() {
	// การตั้งชื่อด้วย camel case เป็น convention ของ Go
	// ถ้า declare variable แล้วไม่ใช้จะ error หรือ import package มาแล้วไม่ใช้ก็จะ error (เพื่อให้ code clean)
	var conferenceName = "Go Conference" // var เก็บค่าสามารถ เปลี่ยนแปลงได้
	const conferenceTickets = 50         // const ไม่สามารถเปลี่ยนได้
	var remainingTickets = 50
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Println("Get your tickets here to attend")
	fmt.Println("=================================")

	// Printf
	// %v => placeholder default format ที่เป็นได้ทั้ง text,number
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("=================================")

	// ทุกครั้งที่ declare variable ต้องบอก data type ด้วย เพื่อให้ Go compiler รู้
	// ถ้าเรา define value ให้ variable Go สามารถรู้ data type ได้โดยที่เราไม่ต้อง define data type (จะ define หรือไม่ define data type ก็ได้)
	// var conferenceName = "Go Conference" <-> var conferenceName string = "Go Conference"
	var userName string
	var userTickets int

	userName = "Gunt"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
	fmt.Println("=================================")

	// print type
	// %T => placeholder ของ type
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)
	fmt.Println("=================================")
	// type
	// int = whole number
	// uint(unsigned integer) = whole number is positive
	// float = suitable for Statistical data,Monetary data

	testValue := "2morrowboyz" // declare,assign value (Syntactic Sugar)
	fmt.Println(testValue)
}
