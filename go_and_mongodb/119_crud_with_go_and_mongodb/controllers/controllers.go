package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"web-dev-go/go_and_mongodb/organizing_code_into_packages/models"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// GetUser attempts to find a user based on a passed in Identification Document "ID"
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	u := models.User{}
	// fetch user
	if err := uc.session.DB("go_web_dev_db").C("users").FindId(oid).One(&u); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		log.Println(err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%s\n", uj)
}

// CreateUser makes a new user
func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from a stream
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// create bson ID
	u.ID = bson.NewObjectId()

	// store the user in mongodb
	err = uc.session.DB("go_web_dev_db").C("users").Insert(u)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
func (uc *UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	oid := bson.ObjectIdHex(id)
	// fetch user
	if err := uc.session.DB("go_web_dev_db").C("users").RemoveId(oid); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	fmt.Fprint(w, "deleted user:", oid)
}
