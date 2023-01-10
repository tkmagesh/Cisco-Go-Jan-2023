package main

import "fmt"

func main() {
	// var nos []int
	nos := make([]int, 0, 5)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 3)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 1)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 4)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 2)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 5)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 8)
	fmt.Printf("len(nos) = %d, cap(nos) = %d, nos = %v\n", len(nos), cap(nos), nos)

	dupNos := nos[2:4]
	fmt.Printf("len(dupNos) = %d, cap(dupNos) = %d, dupNos = %v\n", len(dupNos), cap(dupNos), dupNos)

	dupNos[0] = 100
	fmt.Println("nos :", nos)
	fmt.Println("dupNos :", dupNos)

	x := make([]int, len(nos), len(nos))
	copy(x, nos)
	x[0] = 1000
	fmt.Println("nos :", nos)
	fmt.Println("x :", x)
}
