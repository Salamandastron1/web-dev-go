package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err = sql.Open("mysql", dbConnectString)
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		// exponential retry db connection
		go retryDB(err)
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := io.WriteString(w, "successfully comleted")
	if err != nil {
		log.Println(err.Error())
	}
}

func amigos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos`)
	if err != nil {
		log.Println(err)
	}
	// data to be used in query
	var s, name string
	s = "RETRIEVED RECORDS:\n"

	// query
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Println(err)
		}
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rt, err := stmt.Exec()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	n, err := rt.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "CREATE TABLE customer", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("Tim")`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rt, err := stmt.Exec()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	n, err := rt.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "INSERTED RECORD", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer`)
	if err != nil {
		log.Println(err)
	}
	// data to be used in query
	var name string
	// query
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(name)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Jimmy"`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rt, err := stmt.Exec()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	n, err := rt.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "UPDATED RECORD:", n)
}

func delete(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy"`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rt, err := stmt.Exec()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	n, err := rt.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "UPDATED RECORD:", n)
}

func drop(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Jimmy"`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "DROPPED TABLE customer")
}

func retryDB(err error) {
	count := time.Duration(1 * time.Millisecond)
	for {
		log.Println("DB not ready, retrying...")
		err := db.Ping()
		if err == nil {
			log.Println("DB connection established")
			return
		} else {
			log.Println(err)
			time.Sleep(count)
			count = time.Duration(math.Pow(float64(count+1), float64(count)))
		}
	}
}
