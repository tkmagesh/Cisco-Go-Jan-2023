package main

import "fmt"

func main() {
	// var nos [5]int
	// var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}
	// var nos = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}

	fmt.Println("Iterating using the indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println("Iterating using the range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	fmt.Printf("&nos :%p\n", &nos)
	fmt.Println("After sorting....")
	sort(&nos)
	fmt.Println(nos)

	list1 := [3]string{"Magesh", "Suresh", "Ganesh"}
	list2 := [3]string{"Magesh", "Suresh", "Ganesh"}
	fmt.Printf("%p, %p\n", &list1, &list2)
	fmt.Println(list1 == list2)
}

func sort(list *[5]int) {
	fmt.Printf("list : %p\n", list)
	for i := 0; i <= 3; i++ {
		for j := i + 1; j <= 4; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	fmt.Println("sort completed : ", list)
}
