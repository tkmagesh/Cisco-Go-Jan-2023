package main

import (
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*

Write the apis for the following
Sort => Sort the products collection by any attribute
	IMPORTANT : MUST Use sort.Sort() function
	use cases:
		1. sort the products collection by cost
		2. sort the products collection by name
		3. sort the products collection by units
		etc

*/
type Products []Product

func (products Products) Sort(attrName string) {
	switch attrName {
	case "Id":
		sort.Sort(products)
	case "Name":
		sort.Sort(ByName{products})
	case "Cost":
		sort.Sort(ByCost{products})
	default:
		sort.Sort(products)
	}

}

/* sort.Interface implementation */
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

/* fmt.Stringer interface implementation */
func (products Products) String() string {
	var b strings.Builder
	for _, p := range products {
		b.WriteString(fmt.Sprintf("%s\n", p))
	}
	return b.String()
}

/* to sort by Name */
type ByName struct {
	Products
}

/* overriding the Products.Less() method */
func (this ByName) Less(i, j int) bool {
	return this.Products[i].Name < this.Products[j].Name
}

/* to sort by Cost */
type ByCost struct {
	Products
}

/* overriding the Products.Less() method */
func (this ByCost) Less(i, j int) bool {
	return this.Products[i].Cost < this.Products[j].Cost
}

/* alternate implementation of Sorting */
func (this Products) SortSlice(attrName string) {
	switch attrName {
	case "Id":
		sort.Slice(this, func(i, j int) bool {
			return this[i].Id < this[j].Id
		})
	case "Name":
		sort.Slice(this, func(i, j int) bool {
			return this[i].Name < this[j].Name
		})
	case "Cost":
		sort.Slice(this, func(i, j int) bool {
			return this[i].Cost < this[j].Cost
		})
	case "Units":
		sort.Slice(this, func(i, j int) bool {
			return this[i].Units < this[j].Units
		})
	case "Category":
		sort.Slice(this, func(i, j int) bool {
			return this[i].Category < this[j].Category
		})
	}
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}
	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Default Sort [by id]")
	// products.Sort("Id")
	products.SortSlice("Id")
	fmt.Println(products)

	fmt.Println("Sort by Name")
	// sort.Sort(ByName{products})
	// products.Sort("Name")
	products.SortSlice("Name")
	fmt.Println(products)

	fmt.Println("Sort by Cost")
	// sort.Sort(ByCost{products})
	// products.Sort("Cost")
	products.SortSlice("Cost")
	fmt.Println(products)

	fmt.Println("Sort by Units")
	products.SortSlice("Units")
	fmt.Println(products)

	fmt.Println("Sort by Category")
	products.SortSlice("Category")
	fmt.Println(products)

}
