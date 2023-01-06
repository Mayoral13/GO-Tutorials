/*
Write a program that prints the numbers from 1 to 100, but for multiples of three
print “Fizz” instead of the number and for the multiples of five print “Buzz”. For
numbers which are multiples of both three and five print “FizzBuzz”. (hint: use a
switch with conditions)
*/
package main
import("fmt")
const (
    Fizz = 3;
	Buzz = 5;
	FizzBuzz = 15;
)
func ifc(){
	for i:=0; i<=100; i++ {
		if(i%Fizz == 0){
			fmt.Printf("Fizz: %d\n",i)
		}
		if(i%Buzz == 0){
			fmt.Printf("Buzz: %d\n",i)
		}
		if(i%FizzBuzz == 0){
			fmt.Printf("FizzBuzz: %d\n",i)	
		}else{
			fmt.Printf("%d\n",i)
		}
	}
}
func swi(){
	for i:=0; i<=100; i++ {
		switch{
		case i%Fizz == 0:
			fmt.Printf("Fizz: %d\n",i)
		case i%Buzz == 0:
			fmt.Printf("Buzz: %d\n",i)
		case i%FizzBuzz == 0:
			fmt.Printf("FizzBuzz: %d\n",i)
		default:fmt.Printf("%d\n",i)
		
		}
	}
}
func main(){
	//USED BOTH SWITCH AND IF TO EXPERIMENT AND IF SEEMS TO BE THE BETTER CHOICE
	swi(); //FIZZBUZZ DOESNOT WORK 
	ifc(); //WORKS AS EXPECTED
}