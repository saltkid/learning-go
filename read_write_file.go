package main

import (
	"fmt"
	"os"
)

func write_to_file(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, write_err := file.WriteString("hello file")
	if write_err != nil {
		panic(write_err)
	}

}

func write_lines_to_file(filename string, data []string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, dat := range data {
		_, write_err := fmt.Fprintln(file, dat)
		if write_err != nil {
			panic(write_err)
		}
	}
}

func ReadWriteFile() {
	data := []string{
		"hello",
		"file",
	}
	write_to_file("test_write_file.txt")
	write_lines_to_file("test_write_lines_to_file.txt", data)
}