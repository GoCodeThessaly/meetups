package main

import "fmt"

func main() {
	greetings := make(chan string, 2)

	go func() {
		greetings <- "Hello"
		greetings <- "World!"
	}()

	greet1 := <-greetings
	greet2 := <-greetings
	fmt.Println(greet1, greet2)
}
