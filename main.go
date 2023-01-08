package main

import "fmt"

const engrishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == ""{
		name = "World"
	}
    return engrishHelloPrefix  + name
}

func main() {
    fmt.Println(Hello("world"))
}