package main

type Booking struct {
	Id      int    `json:"id"`
	Name    string `json:"name"` // user name
	ClassId int    `json:"class_id"`
	Date    string `json:"date"`
}
