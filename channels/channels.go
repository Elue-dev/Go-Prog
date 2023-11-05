package main

import (
	"fmt"

	"github.com.elue-dev/go-program/helpers"
)

const NUMBER_POOL = 10

func Calculatevalue(intChann chan int) {
	randomNumber := helpers.RandomNumber(NUMBER_POOL)
	intChann <- randomNumber
} 

func main() {
	intChann := make(chan int)
	defer close(intChann)

	go Calculatevalue(intChann)

	num := <- intChann

	fmt.Println(num)

}



