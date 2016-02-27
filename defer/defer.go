package main

import "fmt"

func main() {
	defer fmt.Println("<3")
	defer fmt.Println("world")	
	fmt.Println("hello")
}
