package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100)
	school := "Holberton School"
	beautifulWeather := true
	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}

	if randomNumber > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
	} else {
		fmt.Printf("my random number is %d and is less than 50\n", randomNumber)
	}

	if school == "Holberton School" {
		fmt.Printf("I go to %s\n", school)
	} else {
		fmt.Printf("Why %s? Everyone should go to Holberton School.\n", school)
	}

	if beautifulWeather {
		fmt.Println("The weather is beautiful today.")
	} else {
		fmt.Println("The weather isn't so beautiful today.")
	}

	for _, value := range holbertonFounders {
		fmt.Printf("%s is a founder\n", value)
	}
}
