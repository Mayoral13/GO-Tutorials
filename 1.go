
/*
Write a program that finds the smallest number
in this list:
x := []int{
 48,96,86,68,
 57,82,63,70,
 37,34,83,27,
 19,97, 9,17,
}
*/
package main
import("fmt")
func main(){
	var apple int;
	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97,9,17,
	   }
	   for i:=0; i<len(x)-1; i++{
		for j:=0; j<len(x); j++{
        if(x[i] < x[j]){
			apple = x[i]
		}else{
			apple = x[j]
		}
		}
	   }
	   fmt.Printf("Smallest is %v\n",apple)
}

