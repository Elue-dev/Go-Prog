package main

import "fmt"

func main() {
	fmt.Println("Hello from Go.")

	whatToSay := "Goodbye from Go"
	var i int
	i = 20
	fmt.Println(whatToSay)
	fmt.Println("i is set to", i)

	whatWasSaid , theOtherThingSaid := saySomething()
	fmt.Println("The function returns", whatWasSaid, theOtherThingSaid)
}

func saySomething() (string, string) {
	return "Something", "Else"
}