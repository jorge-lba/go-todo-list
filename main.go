package main

import (
	"encoding/json"
	"fmt"
	"os"

	// "go-todo-list/domain/entity"
	"log"

	"net/http"

	"go-todo-list/domain/ToDo"
	"go-todo-list/domain/ToDoList"
	"go-todo-list/dtos"
	repo "go-todo-list/repositories/implementation"
	"go-todo-list/repositories/inMemory"
	"go-todo-list/storage"
	model "go-todo-list/storage/gorm"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	cwd, _ := os.Getwd()

	// load .env file
	err := godotenv.Load(cwd + "/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

type CreatedResponse struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

var db, _ = storage.NewDb()

func CreateToToList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db.AutoMigrate(&model.ToDoList{}, &model.ToDo{})

	var tdl ToDoList.Props
	json.NewDecoder(r.Body).Decode(&tdl)

	fmt.Println(tdl)

	t := ToDoList.Create(tdl)

	tdlRepo := repo.ToDoListRepo{
		DB: db,
	}

	tdlRepo.Save(&t)

	inMemory.ToDoListRepo.Save(&t)

	rs := CreatedResponse{
		Message: "To-Do List created successfully",
		Id:      t.Id(),
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rs)
}

func GetToToList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tdlRepo := repo.ToDoListRepo{
		DB: db,
	}

	id := mux.Vars(r)["id"]
	t := tdlRepo.FindById(id)

	dto := dtos.ToDoListEntityToDTO(&t)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dto)
}

func IncludeToDo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tdlRepo := repo.ToDoListRepo{
		DB: db,
	}

	id := mux.Vars(r)["id"]
	t := tdlRepo.FindById(id)

	var tdB ToDo.ToDo
	json.NewDecoder(r.Body).Decode(&tdB)

	tdId := t.IncludeItem(tdB)

	tdlRepo.Save(&t)

	rs := CreatedResponse{
		Id:      tdId,
		Message: "To-Do created successfully",
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rs)
}

func DoneToDo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tdlRepo := repo.ToDoListRepo{
		DB: db,
	}

	id := mux.Vars(r)["id"]
	t := tdlRepo.FindById(id)

	tdId := mux.Vars(r)["todoId"]
	t.DoneToDo(tdId)

	tdlRepo.Save(&t)

	rs := Response{
		Message: "To-Do completed successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rs)
}

func UndoneToDo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tdlRepo := repo.ToDoListRepo{
		DB: db,
	}

	id := mux.Vars(r)["id"]
	t := tdlRepo.FindById(id)

	tdId := mux.Vars(r)["todoId"]
	t.UndoneToDo(tdId)

	tdlRepo.Save(&t)

	rs := Response{
		Message: "To-Do undoned successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rs)
}

func HandleRequest() {
	router := mux.NewRouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./doc/"))).Methods("GET")

	router.HandleFunc("/", CreateToToList).Methods("POST")
	router.HandleFunc("/{id}", GetToToList).Methods("GET")
	router.HandleFunc("/{id}/todo", IncludeToDo).Methods("POST")
	router.HandleFunc("/{id}/todo/{todoId}/done", DoneToDo).Methods("PATCH")
	router.HandleFunc("/{id}/todo/{todoId}/undone", UndoneToDo).Methods("PATCH")

	PORT := goDotEnvVariable("PORT")

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}

func main() {
	HandleRequest()
}
