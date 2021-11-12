package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"web-dev-go/go_and_mongodb/120_thru_122_hands_on_exercises/120/starting-code/models"
	"web-dev-go/go_and_mongodb/120_thru_122_hands_on_exercises/120/starting-code/session"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	session *session.Session
}

func NewUserController(s *session.Session) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := uuid.MustParse(p.ByName("id"))
	u, err := uc.session.GetUser(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = uc.session.CreateUser(&u)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	uj, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := uuid.MustParse(p.ByName("id"))
	// check user and fetch user if available
	err := uc.session.DeleteUser(id)
	if err != nil {
		http.Error(w, "An error has occurred", http.StatusInternalServerError)
		log.Println("Unable to write to file", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user: ", id, "\n")
}
