package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	uuid "github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var xs []string
	checkSession(w, r)
	if r.Method == http.MethodPost {
		xs = addImage(w, r)
	} else {
		c, err := r.Cookie("photos")
		if err != http.ErrNoCookie {
			xs = strings.Split(c.Value, "|")
		}
	}
	if len(xs) > 0 {
		xs = xs[:len(xs)-1]
	}
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func addImage(w http.ResponseWriter, r *http.Request) []string {
	c, err := r.Cookie("photos")
	if err == http.ErrNoCookie {
		c = &http.Cookie{
			Name: "photos",
		}
	}
	fname := createFile(w, r)

	fmt.Println(fname)
	if !strings.Contains(c.Value, fname) {
		c.Value += fmt.Sprintf("%s|", fname)
	}
	http.SetCookie(w, c)

	return strings.Split(c.Value, "|")
}

// createFile creates and copies over the contents of a file
// it returns the new file name stored on the server
func createFile(w http.ResponseWriter, r *http.Request) string {
	mf, fh, err := r.FormFile("nf")
	if err != nil {
		fmt.Println(err)
	}
	defer mf.Close()
	// create sha for file name
	ext := strings.Split(fh.Filename, ".")[1]
	h := sha1.New()
	io.Copy(h, mf)
	fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
	// create new file
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	path := filepath.Join(wd, "public", "pics", fname)
	nf, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer nf.Close()
	//copy
	mf.Seek(0, 0)
	io.Copy(nf, mf)

	return fname
}

func checkSession(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Check Session")
	_, err := r.Cookie("session")
	if err == http.ErrNoCookie {
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: uuid.NewV4().String(),
		})
	}
}
