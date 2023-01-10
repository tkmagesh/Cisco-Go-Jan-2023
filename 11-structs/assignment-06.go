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

func main() {
	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Category: "Stationary",
	}
	fmt.Println("Before Applying discount :", FormatProduct(pen))
	ApplyDiscount(&pen, 10)
	fmt.Println("After Applying discount :", FormatProduct(pen))
}

func FormatProduct(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Category = %q", product.Id, product.Name, product.Cost, product.Category)
}

func ApplyDiscount(product *Product, discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}
