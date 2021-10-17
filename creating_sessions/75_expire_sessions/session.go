package main

import (
	"fmt"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
	}
	// reset session time
	c.MaxAge = sessionLength
	http.SetCookie(w, c)

	// if the user exists, get user
	var u user
	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.username]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	session := dbSessions[c.Value]
	_, ok := dbUsers[session.username]

	return ok
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN")
	showSessions() // for demonstration purposes
	for k, v := range dbSessions {
		if time.Since(v.lastActivity) > time.Second*30 || v.username == "" {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	showSessions() // demonstration purposes
}

func showSessions() {
	fmt.Println("****************")
	for k, v := range dbSessions {
		fmt.Println(k, v.username)
	}
	fmt.Println()
}
