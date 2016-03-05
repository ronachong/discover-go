package main

import (
	"fmt"
	"time"
)

func main() {
	dt := time.Now()
	y := time.Now().Year()
	m := time.Now().Month()
	dy := time.Now().Day()

	fmt.Println("Hello, we are Holberton School")
	fmt.Printf("the date is %s\n", dt)
	fmt.Printf("the year is %d\n", y)
	fmt.Printf("the month is %s\n", m)
	fmt.Printf("the day is %d\n", dy)
}
