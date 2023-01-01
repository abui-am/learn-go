package main

import "fmt"

const NMAX = 5

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
	var employees Employees

	fillArray(&employees)
	fmt.Print("hello world")

	for i := 0; i < len(employees); i++ {
		fmt.Print(employees[i])
	}
}

func fillArray(employees *Employees) {
	for i := 0; i < NMAX; i++ {
		e := &employees[i]
		fmt.Scan(e.id, e.firstName, e.lastName, e.phoneNumber, e.email, e.division, e.branch)
	}
}
