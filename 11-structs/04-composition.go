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

type Dummy struct {
	Name string
}

type PerishableProduct struct {
	Product // -> composition (aka inheritence)
	Expiry  string
}

func NewPerishableProduct(id int, name string, cost float32, category string, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Category: category,
		},
		Expiry: expiry,
	}
}
func main() {
	// grapes := PerishableProduct{Product{101, "Grapes", 50, "Fruits"}, "3 Days"}
	/*
		grapes := PerishableProduct{
			Product: Product{
				Id:       101,
				Name:     "Grapes",
				Cost:     50,
				Category: "Fruits"},
			Expiry: "3 Days",
		}
	*/
	grapes := NewPerishableProduct(101, "Grapes", 50, "Fruits", "3 Days")
	fmt.Println(grapes.Expiry)
	fmt.Println(grapes.Product.Name)
	fmt.Println(grapes.Name)

	fmt.Println(FormatProduct(grapes.Product))
	ApplyDiscount(&grapes.Product, 10)
	fmt.Println(FormatProduct(grapes.Product))
}

func FormatProduct(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %.2f, Category = %q", product.Id, product.Name, product.Cost, product.Category)
}

func ApplyDiscount(product *Product, discount float32) {
	product.Cost = product.Cost * ((100 - discount) / 100)
}
