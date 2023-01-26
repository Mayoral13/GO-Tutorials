/*
5. Calculator
The calculator is one of the most useful and essential tools available on most computers
and smartphones. It’s also a great way to help you understand the basics of working
with code
*/
package main

import("fmt"
	
)
func Add(a...int){
	var value int;
for i := 0; i< len(a); i++{
value += a[i];
}
fmt.Println(value)	
}
func Division(a,b int){
	value := a/b;
	fmt.Println(value)
}
func Multiplication(a,b int){
	value := a*b;
	fmt.Println(value)
}
func Subtraction(a,b int){
		value := a-b;
		fmt.Println(value)
	}
	func main(){

Add(34,55,66,87,88,888,88887)
Division(50,10)
Multiplication(3,5)
Subtraction(100,65)
	}	
