/*
Create a "Product" struct with the following fields
	Id
	Name
	Cost
	Category

Write the following functions
	FormatProduct -> return a formatted string with all the field values of the given product
	ApplyDiscount -> applies the given discount in percentage (update the cost) on the given product
*/

package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Category string
}

func (this Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Category = %q", this.Id, this.Name, this.Cost, this.Category)
}

func (this *Product) ApplyDiscount(discount float32) {
	this.Cost = this.Cost * ((100 - discount) / 100)
}

func main() {
	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Category: "Stationary",
	}
	// fmt.Println("Before Applying discount :", FormatProduct(pen))
	fmt.Println("Before Applying discount :", pen.Format())
	// (&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10)
	/* ApplyDiscount(&pen, 10)
	fmt.Println("After Applying discount :", FormatProduct(pen)) */
	fmt.Println("After Applying discount :", pen.Format())
}
