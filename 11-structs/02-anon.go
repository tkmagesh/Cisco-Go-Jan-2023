package main

import "fmt"

func main() {
	emp := struct {
		id     int
		name   string
		salary float32
	}{
		id:     100,
		name:   "Magesh",
		salary: 10000,
	}
	// fmt.Println(emp)
	// fmt.Printf("%#v\n", emp)
	// fmt.Printf("%+v\n", emp)
	printEmployee(emp)
}

func printEmployee(emp struct {
	id     int
	name   string
	salary float32
}) {
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", emp.id, emp.name, emp.salary)
}
