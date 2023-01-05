package main
 var a = "G"
 func main() {
 n()//PRINTS G
 m()//PRINTS O
 n()//PRINTS O
 }
 func n() {
 print(a)
 }
 func m() {
 a = "O"
 print(a)
 }