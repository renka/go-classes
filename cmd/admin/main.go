package main

// Entry point
var Classes []SingleClass

func main() {
	Classes = []SingleClass{
		{Class: Class{Id: 1, Name: "Pilates", Capacity: 20}, Date: "2020-07-09"},
		{Class: Class{Id: 2, Name: "Yoga", Capacity: 10}, Date: "2020-07-10"},
	}
	handleRequests()
}
