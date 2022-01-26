package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

type Book struct {
	ISBN   string
	Title  string
	Author string
	Price  float32
}

var db *sql.DB
var tpl *template.Template

func init() {
	var err error
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// docker run -p 5432:5432 --name some-postgres -e POSTGRES_PASSWORD=password -d postgres
	db, err = sql.Open("postgres", "postgres://bond:password@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	tpl = template.Must(template.ParseGlob("templates/*"))
	fmt.Println("You connected to the DB")
}

func main() {
	http.HandleFunc("/books", booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/create", booksCreateForm)
	http.HandleFunc("/books/create/process", booksCreateProcess)
	http.HandleFunc("/books/update", booksUpdateForm)
	http.HandleFunc("/books/update/process", booksUpdateProcess)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("method not allowed, request denied")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT * FROM books;")
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.ISBN, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		bks = append(bks, bk)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, $%.2f\n", bk.ISBN, bk.Title, bk.Price)
	}
}

func booksShow(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("method not allowed, request denied")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	bk := Book{}
	err := row.Scan(&bk.ISBN, &bk.Title, &bk.Author, &bk.Price)
	switch {
	case err == sql.ErrNoRows:
		log.Println("Resource not found")
		http.NotFound(w, r)
		return
	case err != nil:
		log.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s, %s, $%.2f\n", bk.ISBN, bk.Title, bk.Price)
}

func booksCreateForm(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "create.gohtml", nil)
}
func booksCreateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("method not allowed, request denied")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	bk := Book{}
	bk.ISBN = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	if bk.ISBN == "" || bk.Title == "" || bk.Author == "" {
		log.Println("missing field")
		http.Error(w, http.StatusText(400)+"missing field", http.StatusBadRequest)
		return
	}
	p, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400)+err.Error(), http.StatusBadRequest)
		return
	}
	bk.Price = float32(p)
	_, err = db.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", bk.ISBN, bk.Title, bk.Author, bk.Price)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400)+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(bk)
	tpl.ExecuteTemplate(w, "created.gohtml", bk)
}

func booksUpdateForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("method not allowed, request denied")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		http.Error(w, http.StatusText(400)+" missing ISBN for book", http.StatusBadRequest)
		return
	}
	row := db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	bk := Book{}
	err := row.Scan(&bk.ISBN, &bk.Title, &bk.Author, &bk.Price)
	switch {
	case err == sql.ErrNoRows:
		log.Println("Resource not found")
		http.NotFound(w, r)
		return
	case err != nil:
		log.Println(err)
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	fmt.Println("i made it to execute template on line 176")
	fmt.Println(bk)
	err = tpl.ExecuteTemplate(w, "update.gohtml", bk)
	if err != nil {
		log.Println(err)
	}
}

func booksUpdateProcess(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Println("method not allowed, request denied")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	bk := Book{}
	bk.ISBN = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	if bk.ISBN == "" || bk.Title == "" || bk.Author == "" {
		log.Println("missing field")
		http.Error(w, http.StatusText(400)+"missing field", http.StatusBadRequest)
		return
	}
	p, err := strconv.ParseFloat(r.FormValue("price"), 32)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400)+err.Error(), http.StatusBadRequest)
		return
	}
	bk.Price = float32(p)
	_, err = db.Exec("INSERT INTO books (title, author, price) VALUES ($1, $2, $3, $4)", bk.ISBN, bk.Title, bk.Author, bk.Price)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(400)+err.Error(), http.StatusBadRequest)
		return
	}
	err = tpl.ExecuteTemplate(w, "updated.gohtml", bk)
	if err != nil {
		log.Println(err)
	}
}
