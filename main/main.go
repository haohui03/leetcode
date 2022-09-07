package main

import "fmt"

func main() {
	output := [5][5]int{}
	fmt.Printf("%T", output)
	fmt.Printf("%T", output[0])
}
