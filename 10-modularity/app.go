package main

import (
	"fmt"

	"github.com/fatih/color"
	calc "github.com/tkmagesh/cisco-go-jan-2023/10-modularity/calculator" /* package alias */
	"github.com/tkmagesh/cisco-go-jan-2023/10-modularity/calculator/utils"
	//_ "github.com/tkmagesh/cisco-go-jan-2023/10-modularity/calculator" /* importing ONLY to execute the init functions */
)

func main() {
	color.Red("app invoked")
	fmt.Println(greet("Magesh"))
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Op Count :", calc.OpCount())
	fmt.Println("Is 21 an even number? :", utils.IsEven(21))
}
