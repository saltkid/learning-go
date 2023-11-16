package main

import (
	"fmt"
	"time"
)

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
	birthday string
}
func (h Hololive) age() int {
	birthday_datetime, err := time.Parse("2006-01-02", h.birthday)
	if err != nil {
		panic(err)
	}

	bday_year := birthday_datetime.Year()
	bday_month := birthday_datetime.Month()
	bday_day := birthday_datetime.Day()

	current_year := time.Now().Year()
	current_month := time.Now().Month()
	current_day := time.Now().Day()

	birthday_occured := (current_month > bday_month) || (current_month == bday_month && current_day >= bday_day)

	age := int(current_year - bday_year)
	if birthday_occured == false {
		age -= 1
	}
	return age
}
func (h Hololive) next_birthday() string {
	birthday_datetime, err := time.Parse("2006-01-02", h.birthday)
	if err != nil {
		panic(err)
	}

	current_year := time.Now().Year()
	bday_month := birthday_datetime.Month()
	bday_day := birthday_datetime.Day()

	next_bday_datetime := time.Date(current_year, bday_month, bday_day, 0,0,0,0, time.UTC)
	next_bday := next_bday_datetime.Format("2006-01-02")

	return next_bday
}
func (h Hololive) age_on_next_birthday() int {
	birthday_datetime, err := time.Parse("2006-01-02", h.birthday)
	if err != nil {
		panic(err)
	}

	bday_year := birthday_datetime.Year()
	current_year := time.Now().Year()

	next_age := current_year - bday_year
	return next_age
}
type Fan struct {
	name string
	trait string
	birthday string
}
func (f Fan) age() int {
	birthday_datetime, err := time.Parse("2006-01-02", f.birthday)
	if err != nil {
		panic(err)
	}

	bday_year := birthday_datetime.Year()
	bday_month := birthday_datetime.Month()
	bday_day := birthday_datetime.Day()

	current_year := time.Now().Year()
	current_month := time.Now().Month()
	current_day := time.Now().Day()

	birthday_occured := (current_month > bday_month) || (current_month == bday_month && current_day >= bday_day)

	age := int(current_year - bday_year)
	if birthday_occured == false {
		age -= 1
	}
	return age
}

type Person interface {
	age() int
}

func years_till_30(p Person) int {
	return 30 - p.age()
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

	aqua := Hololive{"aqua", "cute", "2004-12-01"}
	fmt.Println("______________________\n", aqua.name, "is", aqua.trait, "and is", aqua.age(), "years old")
	fmt.Println(aqua.name, "will be", aqua.age_on_next_birthday(), "on her next birthday on", aqua.next_birthday())

	me := Fan{"my name", "my trait", "2003-12-01"} // all fake if you didn't notice
	people := []Person {
		aqua,
		me,
	}

	for i,person := range people {
		fmt.Print(i+1, ". you have ", years_till_30(person), " years till you're 30 ", person, "\n")
	}

	entry()
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
	fmt.Println("__________________________________\n", dynamic_slice)

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
	fmt.Println("________________")
	fmt.Println("aqua is", my_map["aqua"])
	fmt.Println("shion is", my_map["shion"])
	fmt.Println("ojou is", my_map["ojou"])

	value, exists := my_map["me"]
	fmt.Println("do i exist?", exists)
	if exists != false {
		fmt.Println("yes i do:", value)
	}
}
