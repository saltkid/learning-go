package main

import (
	"fmt"
	"io/fs"
	"os"
)

func write_to_file(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, write_err := fmt.Fprintln(file, "hello file")
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

func append_to_file(filename string, data string) {
	file, err := os.OpenFile(filename,
							os.O_APPEND|os.O_WRONLY,
							fs.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, append_error := fmt.Fprintln(file, data)
	if append_error != nil {
		panic(append_error)
	}
}

func append_lines_to_file(filename string, data []string) {
	file, err := os.OpenFile(filename,
							os.O_APPEND|os.O_WRONLY,
							fs.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, line := range data {
		_, append_error := fmt.Fprintln(file, line)
		if append_error != nil {
			panic(append_error)
		}
	}
}

func ReadWriteFile() {
	// data := []string{
	// 	"hello",
	// 	"file",
	// }
	// write_to_file("test_write_file.txt")
	// write_lines_to_file("test_write_lines_to_file.txt", data)
	
	append_to_file("test_write_file.txt", "i am a new line that was appended")

	append_data := []string {
		"append",
		"THIS",
	}
	append_lines_to_file("test_write_lines_to_file.txt", append_data)
}