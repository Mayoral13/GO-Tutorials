
/*
Write a function with one variadic parameter
that finds the greatest number in a list of numbers.
*/
package main
import "fmt"
func greater(a...int) (int){
	var k int;
	for i:=0; i<len(a) -1 ;i++{
		for j:=0; i<len(a);i++{
		if(a[i] > a[j]){
          k = a[i]
		}else{
			k = a[j]
		}
		}
	}
	return k
}
func main(){
a := greater(1,2,3,5,7,9,990)
fmt.Print("The largest number is: ",a)
}
