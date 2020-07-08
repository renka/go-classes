package main

// Entry point
var Bookings []Booking

func main() {
	Bookings = []Booking{
		{Id: 1, Name: "Hanna", ClassId: 1, Date: "2020-07-09"},
		{Id: 2, Name: "Josef", ClassId: 2, Date: "2020-07-10"},
	}
	handleRequests()
}
