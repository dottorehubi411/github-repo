package main

import "fmt"

func main() {
    // Example usage of some functions or variables
    var score int = 90
    greet("Alice")
    printName()
}

func greet(name string) {
    fmt.Println("Hello, " + name)
}

func printName() {
    fmt.Println("Your name is: ", name)
}
