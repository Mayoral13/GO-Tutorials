/*
4. Rock, Paper, Scissors
Rock, paper, scissors is a straightforward game in real life that’s helped solve arguments
and determine who does what. However, coding a game of rock, paper, scissors is more
complex than you might think, making it a great intermediate project for anyone.
The rock, paper, scissors project uses a random number generator, similar to the dice
rolling project. However, the project goes further by requiring you to code in the rules
and logic of the game and create additional choices and elements for a user to pick.
*/
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)
func main(){
	fmt.Print("Welcome to a game of rock paper and scissors\n")
	var userScore int;
	var BotScore int;
	var options = [3]string{"rock","paper","scissors"}
	for (userScore < 5 || BotScore < 5 ) {
	input := bufio.NewReader(os.Stdin)
	fmt.Println("(Rock Paper or Scissors)")
	text,_ :=input.ReadString('\n') 
	pick := strings.TrimSpace(text);
	if(pick == ""){
		fmt.Println("Empty input retry")
	}else if(pick != "rock" && pick != "paper" && pick != "scissors"){
    fmt.Println("Invalid Input")
	}else{
	random :=rand.Intn(len(options));
	BotPick := options[random];
	fmt.Println("Bot Pick is : "+ BotPick);
	fmt.Println("Your pick is : "+ pick);
	if BotPick == "paper" && pick == "rock"{
		fmt.Println("You lose!")
		BotScore++;
	   }else if BotPick == "paper" && pick == "scissors"{
		fmt.Println("You win!")
		userScore++;
	   }else if BotPick == "rock" && pick == "scissors"{
		fmt.Println("You lose!")
		BotScore++;
	   }else if BotPick == "rock" && pick == "paper"{
		fmt.Println("You win!")
		userScore++;
	   }else if BotPick == "scissors" && pick == "rock"{
		fmt.Println("You win!")
		userScore++;
	   }else if BotPick == "scissors" && pick == "paper"{
		fmt.Println("You lose!")
		BotScore++;
	   }else{
		fmt.Println("Draw!")
	}
	if(BotScore == 5 && userScore < 5){
		fmt.Println("You LOSE")
		fmt.Println("User Score is : ", userScore)
		fmt.Println("Bot Score is : ", BotScore)
		os.Exit(0)
	}else if(userScore == 5 && BotScore < 5){
		fmt.Println("You WIN")
		fmt.Println("User Score is : ", userScore)
	    fmt.Println("Bot Score is : ", BotScore)
		os.Exit(0)
	}
	
}
}
	

}
	
	

