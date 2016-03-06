package main

import (
	"fmt"
	"time"
)

func (u *user) helloName() {
	fmt.Printf("Hello %s\n", u.Name)
}

func (u *user) computeAge() {
	d, _ := time.Parse("January 2, 2006", u.DOB)
	var a = int(time.Since(d).Hours() / 8760)
	fmt.Printf("%s who was born in %s would be %d years old today.\n", u.Name, u.City, a)
}

type user struct {
	Name string `json:"name"`
	DOB  string `json:"date_of_birth"`
	City string `json:"city"`
}

func main() {

	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	u.helloName()
	u.computeAge()
}
