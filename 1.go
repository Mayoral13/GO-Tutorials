package main
 var a string
 func main() {
 a = "G"
 print(a)//PRINTS G
 f1()//PRINTS O
 }

 func f1() {
 a := "O"
 print(a)
 f2()//PRINTS G
 }
 func f2() {
 print(a)
 }