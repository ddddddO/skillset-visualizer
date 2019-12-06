package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

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
	http.HandleFunc("/put/", putGraphDataHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func fetchGraphDataHandler(w http.ResponseWriter, r *http.Request) {
	// Getリクエスト以外受け付けない
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Print("fetchGraphData request")

	// FIXME: QueryRowへ
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

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(categories))
}

func putGraphDataHandler(w http.ResponseWriter, r *http.Request) {
	// Postリクエスト以外受け付けない
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// FIXME:まとめる
	if strings.Count(r.URL.Path, "/") != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if r.ContentLength == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT id FROM skills WHERE id = $1", id)
	if err := row.Scan(); err == sql.ErrNoRows {
		// FIXME: リソース無い的なステータスへ
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	j := `{"frontend": 3,"backend": 3,"database": 3,"infrastructure": 3,"network": 3,"facilitation": 3}`
	result, err := db.Exec("UPDATE skills SET categories = $1", j)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	affectedCnt, _ := result.RowsAffected()
	if affectedCnt != 1 {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Print("putGraphData request")

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello, PUT!"))
	return
}
