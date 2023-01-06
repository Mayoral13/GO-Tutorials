//Create a simple loop with the for construct. Make it loop 15 times and print out the loop
//counter with the fmt package.

//Rewrite this loop using goto. The keyword for may not be used now.
package main 
import("fmt")
func loop15(){
for i:= 0; i<15; i++{
		fmt.Printf("Counter is: %d\n",i)
	}

}
func main(){
	loop15()
}