//Write a function which accepts 2 integers and returns their sum, product and difference. Make a
//version with unnamed return variables, and a 2nd version with named return variables.
package main
import("fmt")
func Calc(a,b uint){
fmt.Printf("Sum is : %d\n",a+b)
fmt.Printf("Product is : %d\n",a*b)
fmt.Printf("Difference is : %d\n",a-b)
}
func main(){
	Calc(6,5);
	
	
}