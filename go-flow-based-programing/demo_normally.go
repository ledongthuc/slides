package main

import "fmt"

func PrintGreeter(name string) {
	greeting := fmt.Sprintf("Hello, %s!", name)
	fmt.Println(greeting)
}

func main() {
	PrintGreeter("Thuc Le")
	PrintGreeter("Lazada")
}
