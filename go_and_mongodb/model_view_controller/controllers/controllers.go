package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"web-dev-go/go_and_mongodb/organizing_code_into_packages/models"

	"github.com/julienschmidt/httprouter"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// GetUser attempts to find a user based on a passed in Identification Document "ID"
func (*UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "james bond",
		Gender: "male",
		Age:    32,
		ID:     p.ByName("id"),
	}

	uj, err := json.Marshal(u)
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser makes a new user
func (*UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from a stream
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u.ID = "007"

	// marshal/unmarshal for having JSON assigned to variable
	uj, err := json.Marshal(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s", uj)
}

// DeleteUser takes an ID and attempts to find if a user is present with that identifier
// if found that user is removed
func (*UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "write code to delete user")
}
