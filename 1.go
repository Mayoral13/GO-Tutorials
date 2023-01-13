
/*
Write a function which takes an integer and
halves it and returns true if it was even or false
if it was odd. For example half(1) should return
(0, false) and half(2) should return (1,
true).
*/
package main
import "fmt"
func half(a int) (int, bool){
if(a %2 == 0){
	b:= a / 2
	return b,true
}else{
	b:= a / 2
	return b,false;
}
}
func main(){
	a,b := half(78)
	fmt.Print("RESULT: ",a,b)
}

