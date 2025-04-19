package main

import "fmt"

var name string = "Pain"
var age int = 30
var fav_number float32 = 3.14
var fun_fact string = "I love automation and retro gaming!"
var likesGoLang bool = true
var message string = "Hi, I'm " + name + ".\n I'm starting my journey and exited to explore concurrency, API, and more!"

func main() {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Favourite Number:", fav_number)
	fmt.Println("Fun Fact:", fun_fact)
	fmt.Println("Learning GO?", likesGoLang)
	fmt.Println("-------------------------------------")
	fmt.Println("Here's a message from me:\n", message)

}
