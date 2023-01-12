//PROGRAM THAT write a program that converts from Fahrenheit
//into Celsius. (C = (F - 32) * 5/9)
package main
import "fmt"
var input float64;
var temp float64;
func main(){
 Converter();
}
func Converter(){
	fmt.Print("WELCOME TO CELSIUS TEMP CONVERTER\n")
	fmt.Print("Enter a value\n")
	fmt.Scanf("%f",&input)
	temp = (input-32)*5 / 9
	fmt.Printf("The inputed value is: %f \n",input)
	fmt.Print("Converted value is: ",temp)
}


