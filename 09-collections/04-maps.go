package main

import "fmt"

func main() {
	// var productsRanks map[string]int
	// var productsRanks map[string]int = make(map[string]int)
	// var productsRanks = make(map[string]int)
	/*
		productsRanks := make(map[string]int)
		productsRanks["pen"] = 4
	*/

	// productsRanks := map[string]int{"pen": 4, "pencil": 1, "marker": 3}
	productsRanks := map[string]int{
		"pen":    4,
		"pencil": 1,
		"marker": 3,
	}
	fmt.Println(productsRanks)

	fmt.Println("Accessing value using the key, productRanks[pen] = ", productsRanks["pen"])
	fmt.Println("len(productRanks) :", len(productsRanks))

	fmt.Println("Iterating a map")
	for key, val := range productsRanks {
		fmt.Printf("productRanks[%q] = %d\n", key, val)
	}

	fmt.Println("Adding a new item")
	productsRanks["scribble-pad"] = 2
	fmt.Println(productsRanks)

	fmt.Println("check if key exists")
	keyToCheck := "pen"
	if rank, exists := productsRanks[keyToCheck]; exists {
		fmt.Printf("Rank of %q is %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("Item %q doesnt exist\n", keyToCheck)
	}

	fmt.Println("Removing an item")
	keyToDelete := "stappler"
	delete(productsRanks, keyToDelete)
	fmt.Println(productsRanks)
}
