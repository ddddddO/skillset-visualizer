package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const DBDSN = "user=postgres dbname=skillsets host=localhost port=5432 sslmode=disable"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", DBDSN)
	if err != nil {
		panic(err)
	}
}

func main() {
	log.Print("start api server")

	Run()
}

func Run() {
	http.HandleFunc("/fetch", fetchGraphDataHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func fetchGraphDataHandler(w http.ResponseWriter, r *http.Request) {
	// Getリクエスト以外受け付けない
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	rows, err := db.Query("SELECT * FROM skills")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id         int64
		categories string
	)
	for rows.Next() {
		if err := rows.Scan(&id, &categories); err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, CATEGORIES: %s", id, categories)
	}

	// あとで、構造体に入れるようにしてマーシャル対応
	body := `{"frontend": 1,"backend": 2,"database": 3,"infrastructure": 3,"network": 2,"facilitation": 4}`
	fmt.Fprint(w, body)
}
