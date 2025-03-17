package main

import "fmt"

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(10)
	fmt.Println("Random Number:", randomNumber)
}
