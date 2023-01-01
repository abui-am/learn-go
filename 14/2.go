package main

import "fmt"

const NMAX = 1000

type Employee struct {
	id          int
	firstName   string
	lastName    string
	phoneNumber int
	email       string
	division    string
	branch      string
}

type Employees [NMAX]Employee

func main() {
	fmt.Print("hello world")
}
