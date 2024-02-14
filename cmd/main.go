package main

import (
	"fmt"
	letters_rand "github.com/matt/aciitools/rand_ascii"
)

func main() {
	fmt.Println(letters_rand.RandASCIIString(10, letters_rand.DIG|letters_rand.UPP))
	fmt.Println(string(letters_rand.RandASCIILower()))
	fmt.Println("Lower:", letters_rand.RandASCIILower())
	fmt.Println(letters_rand.RandASCIIString(4, letters_rand.DIG|letters_rand.PUNCT))

}
