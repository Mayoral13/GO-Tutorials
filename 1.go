/*
Create a program that counts the number of bytes and characters (runes) for this string:
 “asSASA ddd dsjkdsjs dk”
Then do the same for this string: “asSASA ddd dsjkdsjsこん dk”
Explain the difference. (hint: use the unicode/utf8 package.)

*/
package main
import(
	"fmt"
	"unicode/utf8"
)
func main(){
	a := "asSASA ddd dsjkdsjs dk";
	b:= "asSASA ddd dsjkdsjsこん dk"
	fmt.Println("Bytes in A is:",len(a))
	fmt.Println("Bytes in B is:",len(b))
	fmt.Println("Runes in A is:",utf8.RuneCountInString(a))
	fmt.Println("Runes in B is:",utf8.RuneCountInString(b))
}