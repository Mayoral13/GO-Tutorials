//Write another program that converts from feet
//into meters. (1 ft = 0.3048 m)
package main
import "fmt"
var feet float64;
var metre float64;
func main(){
FT_M();
}
func FT_M(){
		fmt.Print("Welcome to feet to metres converter\n")
	fmt.Print("Enter a value\n")
	fmt.Scanf("%f\n",&feet);
	metre = feet * 0.3048;
	fmt.Printf("Converted value is: %f\n",metre);
}
