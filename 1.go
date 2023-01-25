/*
2. Creating a dice-rolling simulator is a great beginner project to help you learn. In
the real world, rolling dice allows you to generate a random number ranging from 1
through 6, so your program should have the same functionality. This project uses the
random module to generate random numbers and print them to the user.
*/
package main

import (
	"fmt"
	"math/rand"
)
func main(){
Random()
}
func Random(){
value := rand.Intn(6)
fmt.Print(value)
}
