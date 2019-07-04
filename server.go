package main

import (
	"./connect"
	"./structures"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	connect.InitializeDatabase()
	defer connect.CloseConnection()
	r := mux.NewRouter()
	//Creamos la ruta que es accedida por GET
	r.HandleFunc("/user/{id}", GetUser).Methods("GET")
	r.HandleFunc("/user/new", NewUser).Methods("POST")
	r.HandleFunc("/user/update/{id}", UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/delete/{id}", DeleteUser).Methods("DELETE")



	log.Println("El servidor se encuentra en el puerto 8080")
	log.Fatal(http.ListenAndServe(":8080", r) )
}

func GetUser(w http.ResponseWriter, r* http.Request)  {
	//esto devuelve un texto
	//w.Write([]byte("Taller Rest!\n"))
	//devolver json

	//leer parametro de la URL
	vars := mux.Vars(r)
	user_id := vars["id"]
	fmt.Println(user_id)

	status := "success"
	var message string
	user := connect.GetUser(user_id)

	//validacion que exista usuario
	if(user.Id <= 0){
		status = "error"
		message = "User not found."
	}

	response := structures.Response{ status, user, message}
	//json creado manualmente
	//user := User{"Eduardo Ismael", "Test1" , "Test2"}
	json.NewEncoder(w).Encode(response)
}

func NewUser(w http.ResponseWriter, r* http.Request){
	user := GetUserRequest(r)
	connect.CreateUser(user)
	response := structures.Response{"success" , connect.CreateUser(user), ""}
	json.NewEncoder(w).Encode(response)

}

func UpdateUser(w http.ResponseWriter, r* http.Request)  {
	vars := mux.Vars(r)
	user_id := vars["id"]

	user := GetUserRequest(r)
	response := structures.Response{"success" , connect.UpdateUser(user_id, user), ""}
	json.NewEncoder(w).Encode(response)
}

func DeleteUser(w http.ResponseWriter, r* http.Request) {
	vars := mux.Vars(r)
	user_id := vars["id"]

	var user structures.User
	connect.DeleteUser(user_id)
	response := structures.Response{"success", user , ""}
	json.NewEncoder(w).Encode(response)
}

/*con eso toma y almacena todo los datos que mande el usuario*/
func GetUserRequest(r* http.Request) structures.User{
	var user structures.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil{
		log.Fatal(err)
	}
	return user
}