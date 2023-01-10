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

	e1 := Employee{1, "abc", 1000}
	e2 := Employee{1, "abc", 1000}
	fmt.Printf("%p, %p\n", &e1, &e2)
	fmt.Println(e1 == e2)

	fmt.Println("Accessing the fields")
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", emp.id, emp.name, emp.salary)

	empPtr := &Employee{
		id:     200,
		name:   "Suresh",
		salary: 20000,
	}
	// empPtr->id
	// fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", (*empPtr).id, (*empPtr).name, (*empPtr).salary)
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", empPtr.id, empPtr.name, empPtr.salary)
}

func printEmployee(emp Employee) {
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", emp.id, emp.name, emp.salary)
}
