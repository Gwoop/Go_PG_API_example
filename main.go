package main

import (
	"API_Kursach/Model/Structs"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "Kyrsach"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/data/user", UserGet).Methods("Get")
	r.HandleFunc("/data/models", ModelsGet).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}

func UserGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userst := Structs.Users{}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Ошибка подключения")
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from Users")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&userst.ID_Users, &userst.Login, &userst.Passwor, &userst.Roles_ID)
		json.NewEncoder(w).Encode(&userst)
	}
}

func ModelsGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	model := Structs.Model{}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Ошибка подключения")
		return
	}
	defer db.Close()
	rows, err := db.Query("select * from Model")
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&model.ID_Model, &model.Name_Model, &model.Configurat_ID, &model.Marks_ID)
		json.NewEncoder(w).Encode(&model)
	}
}
