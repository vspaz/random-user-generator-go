package user

import "time"

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Info struct {
	Seed    string `json:"seed"`
	Results int    `json:"results"`
	Page    int    `json:"page"`
	Version string `json:"version"`
}

type Id struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DateOfBirth struct {
	Date time.Time `json:"date"`
	Age  int       `json:"age"`
}

type Login struct {
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	Md5      string `json:"md5"`
	Sha1     string `json:"sha1"`
	Sha256   string `json:"sha256"`
}

type Timezone struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

type User struct {
	Gender      string      `json:"gender"`
	Name        Name        `json:"name"`
	Location    Location    `json:"location"`
	Email       string      `json:"email"`
	Login       Login       `json:"login"`
	DateOfBirth DateOfBirth `json:"dob"`
	Registered  Registered  `json:"registered"`
	Phone       string      `json:"phone"`
	Cell        string      `json:"cell"`
	Id          Id          `json:"id"`
	Nat         string      `json:"nat"`
}

type Street struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type Location struct {
	Street      Street      `json:"street"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	Postcode    string      `json:"postcode"`
	Coordinates Coordinates `json:"coordinates"`
	Timezone    Timezone    `json:"timezone"`
}

type Registered struct {
	Date time.Time `json:"date"`
	Age  int       `json:"age"`
}

type Users struct {
	Results []User `json:"results"`
}
