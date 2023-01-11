package main

import "fmt"

type MyStr string // type alias

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	str := MyStr("Labore mollit ea laboris ex deserunt ad nisi. Commodo ex nulla eiusmod sit irure pariatur tempor nulla labore voluptate proident sit. Et nulla ea reprehenderit aliqua excepteur officia adipisicing. Incididunt dolor ex exercitation id occaecat occaecat do eu laborum dolor ex voluptate velit. Cillum dolor amet aliquip magna nostrud qui non consectetur id exercitation labore quis ut esse. Sint elit amet irure duis in labore anim do labore in elit ea.")
	fmt.Println(str.Length())
}
