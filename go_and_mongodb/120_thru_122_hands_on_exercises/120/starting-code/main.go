package main

import (
	"log"
	"net/http"
	"web-dev-go/go_and_mongodb/120_thru_122_hands_on_exercises/120/starting-code/controllers"
	"web-dev-go/go_and_mongodb/120_thru_122_hands_on_exercises/120/starting-code/session"

	"github.com/julienschmidt/httprouter"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := httprouter.New()
	// Get a UserController instance
	uc := controllers.NewUserController(session.New())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
