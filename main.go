package main

import "fmt"

func main() {
	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())
}

func increment2() int {
	count := 0
	count++
	return count
}
