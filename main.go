package main

import (
	"fmt"
	"math/rand"
)
func main(){
	secretnumber := rand.Intn(10)+1 // for number ranging from 1-10
	hasWon := false

     for i:=1 ; i <= 5; i++ {
	var guess int
	fmt.Println("Enter your guess no between 0-10:")
	fmt.Scan(&guess)

	if guess == secretnumber {
		fmt.Println("Congrats you guessed the number correctly")
		hasWon = true
		break
	} else {
		attemptsLeft := 5-i

		fmt.Println("Wrong Guess. Attempts left:",attemptsLeft)
		
	}
	
	
}
if !hasWon {
	fmt.Printf("Sorry you are out of attempts. Secret number was %d \n",secretnumber)
}
	fmt.Println("Game End")
}