package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

const DBDSN = "user=postgres dbname=skillsets host=localhost port=5432 sslmode=disable"

var db *sql.DB

func init() {
	log.Print("connect db")

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
	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/fetch", fetchGraphDataHandler)
	http.HandleFunc("/put/", putGraphDataHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("health api")
	// GETリクエスト以外受け付けない
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
}

func fetchGraphDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("fetch graph conn")

	// GETリクエスト以外受け付けない
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var (
		id         int64
		categories string
	)
	row := db.QueryRow("SELECT id, categories FROM skills WHERE id = 1")
	if err := row.Scan(&id, &categories); err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	log.Print("fetchGraphData request")
	log.Printf("ID: %d, CATEGORIES: %s", id, categories)

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(categories))
}

func putGraphDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("put graph conn")

	// OPTIONS・PUTリクエスト以外受け付けないガード節
	switch r.Method {
	case http.MethodPut:
	case http.MethodOptions:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// OPTIONSリクエストの場合
	// プリフライトリクエスト: https://developer.mozilla.org/ja/docs/Web/HTTP/CORS
	if r.Method == http.MethodOptions {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "OPTIONS, PUT")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Access-Control-Max-Age", "86400") // 24h
		w.WriteHeader(http.StatusOK)
		return
	}

	// PUTリクエストの場合
	// TODO:まとめる
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
	targetID, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")

	row := db.QueryRow("SELECT id FROM skills WHERE id = $1", targetID)
	if err := row.Scan(); err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	buf := make([]byte, r.ContentLength)
	_, err = r.Body.Read(buf)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("%+v", err)
		return
	}
	defer r.Body.Close()

	/*
		// debug
		dump, _ := httputil.DumpRequest(r, true)
		log.Print(string(dump))
	*/

	result, err := db.Exec("UPDATE skills SET categories = $1 WHERE id = $2", string(buf), targetID)
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
	return
}
