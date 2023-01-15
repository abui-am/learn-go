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
	var editIndex int
	fillArray(&employees)

	fmt.Scan(&editIndex)
	editData(editIndex, &employees)
	sortDataByFirstName(&employees)
	showData(&employees)
}

func fillArray(employees *Employees) {
	for i := 0; i < NMAX; i++ {
		e := &employees[i]
		fmt.Scan(&e.id, &e.firstName, &e.lastName, &e.phoneNumber, &e.email, &e.division, &e.branch)
	}

}

func editData(i int, employees *Employees) {

	if i >= NMAX || i < 0 {
		fmt.Print("Data tidak ditemukan")
	} else {
		var e = &employees[i]
		// *ID must not be editable
		fmt.Scan(&e.firstName, &e.lastName, &e.phoneNumber, &e.email, &e.division, &e.branch)
	}
}

func sortDataByFirstName(t *Employees) {
	for i := 0; i < NMAX; i++ {
		minIndex := i
		for j := i + 1; j < NMAX; j++ {
			if t[j].firstName < t[minIndex].firstName {
				minIndex = j
			}
		}
		// Swap the minimum element with the current element
		t[i], t[minIndex] = t[minIndex], t[i]
	}
}

func showData(employees *Employees) {
	for i := 0; i < NMAX; i++ {
		e := &employees[i]
		fmt.Printf("%d %s %s %d %s %s %s\n", e.id, e.firstName, e.lastName, e.phoneNumber, e.email, e.division, e.branch)
	}

}
