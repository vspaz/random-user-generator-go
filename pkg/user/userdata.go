package user

import "time"

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Info struct {
	Seed    string `json:"seed,omitempty"`
	Results int    `json:"results,omitempty"`
	Page    int    `json:"page,omitempty"`
	Version string `json:"version,omitempty"`
}

type Id struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type DateOfBirth struct {
	Date time.Time `json:"date,omitempty"`
	Age  int       `json:"age,omitempty"`
}

type Login struct {
	Uuid     string `json:"uuid,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Salt     string `json:"salt,omitempty"`
	Md5      string `json:"md5,omitempty"`
	Sha1     string `json:"sha1,omitempty"`
	Sha256   string `json:"sha256,omitempty"`
}

type Timezone struct {
	Offset      string `json:"offset,omitempty"`
	Description string `json:"description,omitempty"`
}

type User struct {
	Gender      string      `json:"gender,omitempty"`
	Name        Name        `json:"name,omitempty"`
	Location    Location    `json:"location,omitempty"`
	Email       string      `json:"email,omitempty"`
	Login       Login       `json:"login,omitempty"`
	DateOfBirth DateOfBirth `json:"dob,omitempty"`
	Registered  Registered  `json:"registered,omitempty"`
	Phone       string      `json:"phone,omitempty"`
	Cell        string      `json:"cell,omitempty"`
	Id          Id          `json:"id,omitempty"`
	Nat         string      `json:"nat,omitempty"`
}

type Street struct {
	Number int    `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Coordinates struct {
	Latitude  string `json:"latitude,omitempty"`
	Longitude string `json:"longitude,omitempty"`
}

type Location struct {
	Street      Street      `json:"street,omitempty"`
	City        string      `json:"city,omitempty"`
	State       string      `json:"state,omitempty"`
	Country     string      `json:"country,omitempty"`
	Postcode    interface{} `json:"postcode,omitempty"`
	Coordinates Coordinates `json:"coordinates,omitempty"`
	Timezone    Timezone    `json:"timezone,omitempty"`
}

type Registered struct {
	Date time.Time `json:"date,omitempty"`
	Age  int       `json:"age,omitempty"`
}

type Users struct {
	Results []User `json:"results"`
}
