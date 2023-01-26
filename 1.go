/*
6. Mad Libs Generator
Mad Libs is a word game that many of us played on road trips with family when we were
younger. This game makes an excellent programming project that lets you work with
strings, variables, and printing.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main(){
	var words[]string;
	for true{
		fmt.Println("Input a word")
		input := bufio.NewReader(os.Stdin);
		value,_ :=input.ReadString('\n')
		text := strings.TrimSpace(value);
		if(text == ""){
			fmt.Println("Cannot store empty string")
		}else if !(text == "cancel" || text ==""){
			words = append(words,text);
		}
		if(text == "cancel"){
			for i:=0; i<len(words);i++{
				fmt.Println("The word is: ",words[i]);
				
			}
			os.Exit(0)
		}
	}
	


}