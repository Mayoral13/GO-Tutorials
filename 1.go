/*
7. To-Do List
Creating a to-do list is a great beginner programming project that can be completed in
just about any programming language. Completing a to-do list in JavaScript is a great
beginner project for anyone interested in web development.
Your to-do list will need to allow users to add, complete, and delete tasks.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)
 type Task struct{
	TaskID string `json:"TaskID"`;
	Taskname string `json:"TaskName"`;
	Completed string `json:"Completed"`;
 }
 var tasks[]Task;
func main(){
	tasks = append(tasks, Task{TaskID:"1",Taskname:"Cook",Completed:"false"});
	mux := mux.NewRouter();
	http.Handle("/",mux)
	mux.HandleFunc("/Add",AddTask).Methods("POST")
	mux.HandleFunc("/View",ViewTask).Methods("GET")
	mux.HandleFunc("/View/{id}",ViewOneTask).Methods("GET")
	mux.HandleFunc("/Update/{id}",UpdateTask).Methods("PUT")
	mux.HandleFunc("/Delete/{id}",DeleteTask).Methods("DELETE")
	mux.HandleFunc("/Complete/{id}",CompleteTask).Methods("PUT")
	fmt.Println("WELCOME TO TODO LIST")
	if err:=http.ListenAndServe(":3000",nil);err!=nil{
		log.Fatal(err)
	}


}
func AddTask(w http.ResponseWriter, r *http.Request){
w.Header().Set("Content-Type","application/json")
var data Task;
json.NewDecoder(r.Body).Decode(&data)
tasks = append(tasks,data)
json.NewEncoder(w).Encode(data)
}
func ViewTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(tasks)
}
func ViewOneTask(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ID := mux.Vars(r);
	for _,task := range tasks{
	if(task.TaskID == ID["id"]){
		json.NewEncoder(w).Encode(task)
	}
	}
}
func UpdateTask(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ID := mux.Vars(r);
	for id,task := range tasks{
	if(task.TaskID == ID["id"]){
		tasks = append(tasks[:id],tasks[id + 1:]... )
		var data Task;
		json.NewDecoder(r.Body).Decode(&data)
		tasks = append(tasks,data)
		json.NewEncoder(w).Encode(data)
	}
	}

}
func CompleteTask(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ID := mux.Vars(r);
	for id,task := range tasks{
	if(task.TaskID == ID["id"] && task.Completed != "completed"){
		tasks = append(tasks[:id],tasks[id + 1:]... )
		tasks = append(tasks,Task{TaskID:task.TaskID,Taskname:task.Taskname,Completed:"completed"})
		json.NewEncoder(w).Encode(tasks)
	}else if(task.TaskID == ID["id"] && task.Completed == "completed"){
      json.NewEncoder(w).Encode("TASK IS COMPLETED")
		}
	}

}
func DeleteTask(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	ID := mux.Vars(r);
	for id,task := range tasks{
	if(task.TaskID == ID["id"]){
		tasks = append(tasks[:id],tasks[id + 1:]... )
	}
	}

}


