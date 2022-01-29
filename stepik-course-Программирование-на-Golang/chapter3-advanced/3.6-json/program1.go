package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Group struct {
	ID,
	Number string
	Year     int
	Students []Students
}
type Students struct {
	LastName,
	FirstName,
	MiddleName,
	Birthday,
	Addres,
	Phone string
	Rating []int
}
type Rating struct {
	Average float32
}

func Program1() {
	f, err := os.Open("data.json")
	if err != nil {
		return
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	if !json.Valid(data) {
		return
	}
	group := &Group{}
	json.Unmarshal(data, group)

	sum := float32(0.0)
	for _, student := range group.Students {
		sum += float32(len(student.Rating))
	}
	rating := Rating{sum / float32(len(group.Students))}
	output, err := json.MarshalIndent(rating, "", "    ")
	if err != nil {
		return
	}
	wr := bufio.NewWriter(os.Stdout)
	wr.WriteString(string(output) + "\n")
	wr.Flush()
}
