package helpers

import "math/rand"

type SomeType struct {
	Name string
	Age  int
}

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}