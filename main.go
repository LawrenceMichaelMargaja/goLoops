package main

import "fmt"

func main() {
	fruits := []string{
		"apples", "bananas", "mangoes",
	}

	for index, value := range fruits {
		fmt.Println("the index -- ", index)
		fmt.Println("the value -- ", value)
	}
}
