package main

import "fmt"

func main() {
	fmt.Println(Hello("John"))
	fmt.Println(Hello("Mary"))
}

func Hello(name string) string {
	return "Hello " + name
}
