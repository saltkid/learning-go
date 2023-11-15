package main

import "fmt"

// constants
const org = "Hololive"

// fake enums
const (
	friend = "Shion"
	lover  = "Okayu"
)

func main() {
	fmt.Println("hello go")

	// variable declarations
	var fake_age string
	fake_age = "secret"
	var name string = "Minato Aqua"
	trait := "cute"

	// printing
	fmt.Println("My oshi's name is", name, "and she is", fake_age, "years old")
	fmt.Printf("She is %v\n", trait)

	age := get_age()
	// function calls
	validate_age(age)

	// switch statement
	switch org {
	case "Hololive":
		fmt.Println("She is an idol")
	case "some other org":
		fmt.Println("She is an idol")
	default:
		fmt.Println("SHE IS AN IDOL")
	}
	fmt.Println("She is friends with", friend, "but", friend, "constantly steals the attention of her lover,", lover)

	// for loops (only loop in the language)
	for i := 1; i < 5; i++ {
		fmt.Print(i, ". AKUAAAAAAAAAA\n")
	}
}

// function declaration
func validate_age(age int) {
	// if elses + converting between float and int type
	if float64(age) > 18.0 {
		fmt.Println("She is of legal age :^)")
	} else if age == 18 {
		fmt.Println("Barely...kinda weird")
	} else {
		fmt.Println("She's a minor dawg, quit")
	}
}

// return type declaration
func get_age() (age int) {
	return 15 + 4
}
