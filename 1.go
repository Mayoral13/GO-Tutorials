//Write a function Season which has as input-parameter a month-number and
//which returns the name of the season to which this month belongs (disregard the day in the month).
package main
import("fmt")
func Season(a uint){
switch a {
case 1:fmt.Printf("Summer\n")
case 2:fmt.Printf("Winter\n")
case 3:fmt.Printf("Autumn\n")
default:fmt.Printf("Just a regular day\n")
}
}
func main(){
Season(1)
Season(2)
Season(3)
Season(4)
}
