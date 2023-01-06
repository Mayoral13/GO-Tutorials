package main
import("fmt";
"runtime")
var prompt = "Enter a digit, e.g. 3 "+ "or %s to quit."
func windows(){
if(runtime.GOOS == "windows"){
	prompt = fmt.Sprintf(prompt,"Ctrl + Z, Enter");
	fmt.Print(prompt)
}
}
func linux(){
	if(runtime.GOOS == "linux"){
	prompt = fmt.Sprintf(prompt,"Ctrl + D, Exit");
	fmt.Print(prompt)
}
}
func main(){
windows()
linux()
}