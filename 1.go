//A CRUD API WITH A SLICE AS DATA STORAGE
package main
import(
	//"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
	//"strconv"
	//"math/rand"

)
type Data struct{
	Name string `json:"name"`
	Occupation string `json:"occupation"`
	Age string `json:"age"`
	Id string `json:"id"`
	User *User `json:"user"`
}

type User struct{
	Name string `json:"name"`
	Level string `json:"level"`
}
func create(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var datax Data;
	_ = json.NewDecoder(r.Body).Decode(&data);
    data = append(data,datax)
	json.NewEncoder(w).Encode(datax)
}
func getAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(data)
}
func getOne(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _,item := range data{
		if(item.Id == params["id"]){
			json.NewEncoder(w).Encode(item)
		}else{
			http.Error(w,"NOT FOUND",http.StatusNotFound);
		}
	}
}
func delete(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	params :=mux.Vars(r)
	for index,item := range data{
		if (item.Id == params["id"]){
        data = append(data[:index],data[index+1:]... )
		break
		}else{
			http.Error(w,"NOT FOUND",http.StatusNotFound);
		}
	}
}
func update(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Conteent/Type","application/json")
	params := mux.Vars(r)
	for index,item := range data{
		if(item.Id == params["id"]){
			data = append(data[:index],data[index + 1:]...)
			var datax Data
			_ = json.NewDecoder(r.Body).Decode(&datax)
			datax.Id = params["id"]
			data = append(data,datax)
			json.NewEncoder(w).Encode(datax)
			return
		}
	}
}

var data []Data
func main(){
	r := mux.NewRouter();
    data = append(data,Data{Name:"John",Occupation:"Doctor",Age:"24",Id:"1",User:&User{Name:"Cena",Level:"100"}})
	r.HandleFunc("/create",create).Methods("POST")
	r.HandleFunc("/update/{id}",update).Methods("PUT")
	r.HandleFunc("/get",getAll).Methods("GET")
	r.HandleFunc("/get/{id}",getOne).Methods("GET")
	r.HandleFunc("/delete/{id}",delete).Methods("DELETE")
	if err:=http.ListenAndServe(":3000",nil);err != nil{
		log.Fatal(err)
	}
}