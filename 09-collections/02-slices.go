package main

import "fmt"

func main() {
	// var nos []int
	// var nos []int = []int{3, 1, 4, 2, 5}
	// var nos []int = []int{3, 1, 4}
	// var nos = []int{3, 1, 4, 2, 5}
	nos := []int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating using the indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating using the range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	fmt.Println("appending new values")
	// nos = append(nos, 100)
	// nos = append(nos, 100, 200, 300)
	hundreds := []int{100, 200, 300}
	nos = append(nos, hundreds...)
	fmt.Println(nos)

	fmt.Println("slicing...")
	fmt.Println("nos[2:5] :", nos[2:5])
	fmt.Println("nos[2:] :", nos[2:])
	fmt.Println("nos[:5] :", nos[:5])

	fmt.Println("After sorting")
	sort(nos)
	fmt.Println(nos)
}

func sort(list []int) {
	fmt.Printf("list : %p\n", list)
	for i := 0; i <= 3; i++ {
		for j := i + 1; j <= 4; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	// fmt.Println("sort completed : ", list)
}
