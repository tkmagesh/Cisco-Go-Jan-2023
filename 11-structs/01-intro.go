package main

import "fmt"

type Employee struct {
	id     int
	name   string
	salary float32
}

func main() {
	// var emp Employee
	// emp := Employee{200, "Suresh", 20000}

	emp := Employee{
		id:     100,
		name:   "Magesh",
		salary: 10000,
	}

	// fmt.Println(emp)
	// fmt.Printf("%#v\n", emp)
	fmt.Printf("%+v\n", emp)
	// printEmployee(emp)
}

func printEmployee(emp Employee) {
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", emp.id, emp.name, emp.salary)
}
