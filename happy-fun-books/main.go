package main

import "fmt"

func main() {
	var title string
	var copies int
	var author string

	title = "For the Love of Go"
	copies = 99
	author = "John Arundel"
	edition := "1st"

	printData(title, copies, author, edition)
}

func printData(value ...any) {
	fmt.Println(value)
}
