package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"web-dev-go/go_and_mongodb/120_thru_122_hands_on_exercises/120/starting-code/controllers"
	"web-dev-go/go_and_mongodb/120_thru_122_hands_on_exercises/120/starting-code/models"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func init() {

}

func main() {
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() map[uuid.UUID]models.User {
	var users map[uuid.UUID]models.User = make(map[uuid.UUID]models.User)
	um, err := os.Open("userMap.json")
	if err != nil {
		log.Println(err.Error())
		log.Println("Proceeding with run, creating non-nil map")
	} else {
		json.NewDecoder(um).Decode(&users)
	}

	return users
}
