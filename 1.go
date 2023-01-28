/*
JUST A RANDOM QUIZ GAME
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
    var words = [1]string{"boy"}
	var lives int = 3;
	var guesses int = 3;
	var enteredGame bool;
func main(){
	fmt.Println("Welcome to a game of hangman")
	fmt.Println("Enter 1 to start game")
	for lives != 0{
	input := bufio.NewReader(os.Stdin)
	value,_ := input.ReadString('\n')
	text := strings.TrimSpace(value);
	if(text == "1"){
		fmt.Println("You have started the game")
		enteredGame = true;
		game()
	}else{
		fmt.Println("INVALID INPUT")
	}
}
}
func game(){
   for enteredGame{
	gameOver()
	fmt.Println("LIVES LEFT ARE : ", lives)
	fmt.Println("GUESSES LEFT ARE : ", guesses)
	fmt.Println("First question is what is a young male?")
	input := bufio.NewReader(os.Stdin)
	value,_ := input.ReadString('\n')
	text := strings.TrimSpace(value);
     if(text == "boy"){
      fmt.Println("WELLDONE")
	  gameOver()
}else{
	fmt.Println("WRONG")
	lives--;
	fmt.Println("Do you want to use a guess ?")
	fmt.Println("Press y for yes n for no")
	input := bufio.NewReader(os.Stdin)
	value,_ := input.ReadString('\n')
	text := strings.TrimSpace(value);
	if(text == "y"){
		guessOver()
		fmt.Println("It rhymes with soil")
		lives++;
		guesses--;
	}else if(text == "n" && lives != 0){
		fmt.Println("Try again")
         game()
	}else if(text == "n"){
         game()
	}else{
		fmt.Println("Invalid input")
		
	}
}
}
} 
func gameOver(){
	if(lives == 0){
		fmt.Println("GAME OVER")
		os.Exit(0)
	}
}
func guessOver(){
	if(guesses == 0){
		fmt.Println("You cannot use any more guesses")
		game()
	}
}