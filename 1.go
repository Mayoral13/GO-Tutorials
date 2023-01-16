// CRUD API
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)
func Create(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","aplication/json")
	var datax Data;
	json.NewDecoder(r.Body).Decode(&datax)
	data = append(data,datax);
	json.NewEncoder(w).Encode(datax);
}
func Get(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","aplication/json")
	json.NewEncoder(w).Encode(data);
}
func Update(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","aplication/json")
	params := mux.Vars(r)
	for index,item := range data{
		if(item.ID == params["id"]){
		data = append(data[:index],data[index+1:]... )
		var dataxx Data;
		json.NewDecoder(r.Body).Decode(&dataxx)
		data = append(data,dataxx)
		json.NewEncoder(w).Encode(data)
		}
		
	}

}
func Delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","aplication/json")
	params := mux.Vars(r)
	for index,item := range data{
		if(item.ID == params["id"]){
			data = append(data[:index],data[index + 1:]... )
			json.NewEncoder(w).Encode(data)
		}
		
	}
	
}
type Data struct{
	ID string `json:"ID"`
	Name string `json:"Name"`
	Occupation string `json:"Occupation"`
	Age string `json:"Age"`
	User *User `json:"User"`
}
type User struct{
	Name string `json:"Name"`
	Level string `json:"Level"`
}
var data []Data;

func main(){
	data = append(data,Data{ID:"1",Name:"Kendrick",Occupation:"Rapper",Age:"24",User:&User{Name:"Lamar",Level:"10"}})
	data = append(data,Data{ID:"2",Name:"Kendrick",Occupation:"Rapper",Age:"24",User:&User{Name:"Lamar",Level:"10"}})
	data = append(data,Data{ID:"3",Name:"Kendrick",Occupation:"Rapper",Age:"24",User:&User{Name:"Lamar",Level:"10"}})

	r := mux.NewRouter();
	r.HandleFunc("/Create",Create).Methods("POST")
	r.HandleFunc("/Get",Get).Methods("GET")
	r.HandleFunc("/UpdKate/{id}",Update).Methods("PUT")
	r.HandleFunc("/Delete/{id}",Delete).Methods("DELETE")
	fmt.Print("SERVER IS LIVE")
	log.Fatal(http.ListenAndServe(":3000",r))
}