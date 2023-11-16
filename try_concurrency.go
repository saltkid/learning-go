package main

import (
	"fmt"
	"time"
)

type Manjigumi struct {
	name string
	age int
}

func printMembers(ch chan []Manjigumi) {
	manji := <-ch // recieve the manji slice from the channel

	for _, member := range manji {
		fmt.Printf("Name: %s, Age: %d\n", member.name, member.age)
	}
}

func entry() {
	// make channel of manjigumi members
	manji_ch := make(chan []Manjigumi)

	var manji []Manjigumi
	manji = append(manji, Manjigumi{"aqua", 18})
	manji = append(manji, Manjigumi{"ojou", 18})
	manji = append(manji, Manjigumi{"shion", 18})

	// start goroutine
	go printMembers(manji_ch)
	// send manji slice to manji channel
	manji_ch <- manji
	// close the channel
	time.Sleep(500 * time.Millisecond) // usually, waitgroups are used but sleep is okay for now (bc I don't understand wgs as a concept lol)
	close(manji_ch)
}