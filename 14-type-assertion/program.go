package main

import "fmt"

func main() {
	// var x interface{} // before go 1.18
	var x any // from go 1.18

	// x = 100
	// x = 99.99
	// x = "Amet id minim cillum mollit ullamco cillum quis eu laborum anim. Enim aliquip nisi est tempor minim aliquip adipisicing ex eu anim. Ut cupidatat esse eu elit magna occaecat sint aute non cupidatat quis sit amet deserunt. Incididunt et in voluptate excepteur pariatur occaecat ut cupidatat in aliquip labore sit amet. Labore magna reprehenderit consequat ad ad nulla consectetur reprehenderit ea."
	// x = true
	// x = struct{}{}
	x = 4 + 5i
	fmt.Println(x)
	// x = true
	// y := x.(int) + 200
	if val, ok := x.(int); ok {
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case float64:
		fmt.Println("x is a float64, x / 10.5 =", val/10.5)
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case bool:
		fmt.Println("x is a bool, !x =", !val)
	case struct{}:
		fmt.Println("x is an anonymous struct")
	default:
		fmt.Println("Unknown type")
	}

}
