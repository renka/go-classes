package main

type Class struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type SingleClass struct {
	Class
	ClassId string `json:"class_id"`
	Date    string `json:"date"`
}

type EntireClass struct {
	Class
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type Error struct {
	Status  int
	Message string
}
