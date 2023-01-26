package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"math/rand"
)
func main(){
	fmt.Println("Welcome to a game of Rock Paper Scissors")
	var userScore int;
	var options = [3]string{"rock","paper","scissors"}
	var botScore int;
	for true{
	fmt.Println("rock, paper or scissors ?")
	input := bufio.NewReader(os.Stdin);
	value,_ := input.ReadString('\n');
	text := strings.TrimSpace(value);
	if !(text =="rock" || text =="paper" || text =="scissors"){
		fmt.Println("INVALID INPUT")
	}else{
	botIndex := rand.Intn(len(options))
	botPick := options[botIndex];
	fmt.Println("You picked: "+ text)
	fmt.Println("Bot picked: "+ botPick)
	fmt.Println("UserScore is: ", userScore)
	fmt.Println("BotScore is: ", botScore)
	if(text == botPick){
		fmt.Println("Tie")
	}else if(text == "rock" && botPick =="paper"){
		fmt.Println("You lose");
		botScore++;
		fmt.Println("UserScore is: ", userScore)
	    fmt.Println("BotScore is: ", botScore)
	}else if(text =="paper" && botPick =="scissors"){
		fmt.Println("You lose");
		botScore++;
		fmt.Println("UserScore is: ", userScore)
	    fmt.Println("BotScore is: ", botScore)
	}else if(text == "scissors" && botPick =="rock"){
		fmt.Println("You lose");
		botScore++;
		fmt.Println("UserScore is: ", userScore)
	    fmt.Println("BotScore is: ", botScore)
	}else if(botPick == "rock" && text =="paper"){
		fmt.Println("You win");
		userScore++;
		fmt.Println("UserScore is: ", userScore)
	    fmt.Println("BotScore is: ", botScore)
	}else if(botPick =="paper" && text =="scissors"){
		fmt.Println("You win");
		userScore++;
		fmt.Println("UserScore is: ", userScore)
	    fmt.Println("BotScore is: ", botScore)
	}else if(botPick == "scissors" && text =="rock"){
		fmt.Println("You win");
		userScore++;
		fmt.Println("UserScore is: ", userScore)
	    fmt.Println("BotScore is: ", botScore)
	}
	if(userScore == 5 && botScore < 5){
		fmt.Println("YOU WIN")
		fmt.Println("UserScore is: ", userScore)
	    fmt.Println("BotScore is: ", botScore)
		os.Exit(0)
	}else if(botScore == 5 && userScore < 5){
		fmt.Println("YOU LOSE")
		fmt.Println("UserScore is: ", userScore)
	    fmt.Println("BotScore is: ", botScore)
		os.Exit(0)
	}
	}
	}


}