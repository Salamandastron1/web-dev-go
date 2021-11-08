package main

import (
	"net/http"
	"web-dev-go/go_and_mongodb/119_crud_with_go_and_mongodb/controllers"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":8080", r)
}

func getSession() *mgo.Session {
	// connect to our mongo database
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}

	return s
}
