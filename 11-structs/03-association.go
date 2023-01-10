package main

import "fmt"

type Organization struct {
	name string
	city string
}

type Employee struct {
	id     int
	name   string
	salary float32
	org    *Organization
}

func main() {
	cisco := Organization{
		name: "Cisco",
		city: "Bengaluru",
	}
	e1 := Employee{
		id:     100,
		name:   "Magesh",
		salary: 10000,
		org:    &cisco,
	}
	fmt.Println(e1)
	fmt.Println(e1.org.name, e1.org.city)

	e2 := Employee{
		id:     102,
		name:   "Suresh",
		salary: 20000,
		org:    &cisco,
	}
	fmt.Println("e1 :", e1)
	fmt.Println("e2 :", e2)

	fmt.Println("Chaning the org city to Pune")
	cisco.city = "Pune"
	fmt.Println("e1.org :", e1.org)
	fmt.Println("e2.org :", e2.org)
}
