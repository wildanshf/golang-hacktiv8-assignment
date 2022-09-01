package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	id     int64
	name   string
	street string
	job    string
	quotes string
}

func main() {
	var allStudents = []student{
		{
			id:     1,
			name:   "John",
			street: "St. John",
			job:    "Manager",
			quotes: "I am a quote",
		},
		{
			id:     2,
			name:   "Budi",
			street: "St. Budi",
			job:    "Engineer",
			quotes: "I am a quote",
		},
		{
			id:     3,
			name:   "Andi",
			street: "St. Andi",
			job:    "Doctor",
			quotes: "I am a quote",
		},
		{
			id:     4,
			name:   "Ryan",
			street: "St. Ryan",
			job:    "Driver",
			quotes: "I am a quote",
		},
	}

	first, err := strconv.ParseInt(os.Args[1], 10, 64)

	if err != nil {
		fmt.Println("First input parameter must be integer")
		os.Exit(1)
	}

	foundStudent, isNotfound := findStudent(first, allStudents)

	if isNotFound != nil {
		fmt.Println(isNotFound)
	} else {
		fmt.Printf("%#v", foundStudent)
	}
}

func findStudent(search int64, students []student) (student student, err error) {
	for _, student := range students {
		if student.id == search {
			return student, nil
		}
	}
	err = fmt.Errorf("student with id %d not found", search)
	return student, err
}
