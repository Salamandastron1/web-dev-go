package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"web-dev-go/go_and_mongodb/120_thru_122_hands_on_exercises/120/starting-code/models"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type UserController struct {
	session *session
}
type session struct {
	inMem     map[uuid.UUID]models.User
	fileStore *os.File
}

func NewUserController() *UserController {
	return &UserController{getSession()}
}

func getSession() *session {
	var users map[uuid.UUID]models.User = make(map[uuid.UUID]models.User)
	um, err := os.Open("userMap.json")
	if err != nil {
		log.Println(err.Error())
		log.Println("Proceeding with run, using in-memory non-nil map")
		um, err = os.Create("userMap.json")
		if err != nil {
			log.Fatal("Unable to create backup file: ", err.Error())
		}
		log.Println("'userMap.json' created")
	} else {
		json.NewDecoder(um).Decode(&users)
	}
	defer um.Close()

	return &session{users, um}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := uuid.MustParse(p.ByName("id"))
	// check user and fetch user if available
	u, ok := uc.session.inMem[id]
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
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// store the user in map
	// handle case for unique ID colliding
	for {
		if _, ok := uc.session.inMem[u.ID]; !ok {
			uc.session.inMem[u.ID] = u
			break
		} else {
			u.ID = uuid.New()
		}
	}

	uj, _ := json.Marshal(u)
	uc.writeToFile()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := uuid.MustParse(p.ByName("id"))
	// check user and fetch user if available
	_, ok := uc.session.inMem[id]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
		// delete user
	} else {
		delete(uc.session.inMem, id)
	}
	err := uc.writeToFile()
	if err != nil {
		http.Error(w, "An error has occurred", http.StatusInternalServerError)
		log.Println("Unable to write to file", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user: ", id, "\n")
}

func (uc *UserController) writeToFile() error {
	f, err := os.OpenFile("userMap.json", os.O_WRONLY, os.ModeAppend)
	if err != nil {
		f.Close()
		return fmt.Errorf("unable able to write to file: %s", err.Error())
	}
	if err != nil {
		f.Close()
		return err
	}
	err = json.NewEncoder(f).Encode(uc.session.inMem)
	if err != nil {
		f.Close()
		return err
	}

	return f.Close()
}
