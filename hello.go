package main

import "fmt"

// constants
const org = "Hololive"

// fake enums
const (
	friend = "Shion"
	lover  = "Okayu"
)
// structs
type Hololive struct {
	name string
	trait string
}

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

	// function calls
	age := get_age()
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

	slice_manipulation()
	maps()

	aqua := Hololive{"aqua", "cute"}
	fmt.Println(aqua.name, "is", aqua.trait)
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

func slice_manipulation() {
	dynamic_slice := []int{1,2,3}

	// assigning values
	dynamic_slice[0] = 4
	dynamic_slice[1] = 5
	dynamic_slice[2] = 6
	fmt.Println(dynamic_slice)

	// declare slice first before copying
	copy_slice := make([]int, 3)
	copy(copy_slice, dynamic_slice)
	copy_slice = append(dynamic_slice, 7, 8, 9)
	fmt.Println("original:", dynamic_slice, "\ncopy + new elements:", copy_slice)

	fmt.Println(cap(dynamic_slice), cap(copy_slice))
}

func maps() {
	my_map := make(map[string]string)

	// assigning values
	my_map["aqua"] = "cute"
	my_map["shion"] = "cute"
	my_map["ojou"] = "cute"
	my_map["me"] = "cute"

	// sadge
	delete(my_map, "me")

	fmt.Println("aqua is", my_map["aqua"])
	fmt.Println("shion is", my_map["shion"])
	fmt.Println("ojou is", my_map["ojou"])

	value, exists := my_map["me"]
	fmt.Println("do i exist?", exists)
	if exists != false {
		fmt.Println("yes i do:", value)
	}
}