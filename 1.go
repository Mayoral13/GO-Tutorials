package main
import("fmt";
"runtime")
var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."
func main(){
if(runtime.GOOS == "windows"){
	prompt = fmt.Sprintf(prompt,"Ctrl + Z, Enter");
	fmt.Print(prompt)
}else{
	prompt = fmt.Sprintf(prompt,"Ctrl + D, Enter");
	fmt.Print(prompt);

}

 

}