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
	editData(&employees)
	fmt.Print("hello world")
}

func fillArray(employees *Employees) {
	for i := 0; i < NMAX; i++ {
		e := &employees[i]
		fmt.Scan(&e.id, &e.firstName, &e.lastName, &e.phoneNumber, &e.email, &e.division, &e.branch)
	}

}

func editData(employees *Employees) {
	var i int
	fmt.Scan(&i)

	if i >= NMAX || i < 0 {
		fmt.Print("Data tidak ditemukan")
	} else {
		var e = &employees[i]
		// *ID must not be editable
		fmt.Scan(&e.firstName, &e.lastName, &e.phoneNumber, &e.email, &e.division, &e.branch)
	}
}
