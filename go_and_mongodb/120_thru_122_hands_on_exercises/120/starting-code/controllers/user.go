package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-dev-go/go_and_mongodb/120_thru_122_hands_on_exercises/120/starting-code/models"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	session map[uuid.UUID]models.User
}

func NewUserController(s map[uuid.UUID]models.User) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := uuid.MustParse(p.ByName("id"))
	// check user and fetch user if available
	u, ok := uc.session[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{ID: uuid.New()}
	json.NewDecoder(r.Body).Decode(&u)
	// store the user in map
	// handle case for unique ID colliding
	for {
		if _, ok := uc.session[u.ID]; !ok {
			uc.session[u.ID] = u
			break
		} else {
			u.ID = uuid.New()
		}
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := uuid.MustParse(p.ByName("id"))
	// check user and fetch user if available
	_, ok := uc.session[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
		// delete user
	} else {
		delete(uc.session, id)
	}
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user: ", id, "\n")
}
