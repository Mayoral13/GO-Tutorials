/*1. Odd or Even?
This is about as simple as it comes — this program will take a number input by a user
and then tell the user if the number is odd or even.*/
package main
import("fmt")
var input int;
func main(){
	value,_:=fmt.Scan(&input)
	if(input % 2 == 0){
		fmt.Print("NUMBER IS EVEN")
	}else{
		fmt.Print("NUMBER IS ODD \n")
		fmt.Print(value)

}
}