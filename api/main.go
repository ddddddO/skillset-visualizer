package main

import (
	"fmt"
	"log"
	"net/http"
)

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

	// あとで、構造体に入れるようにしてマーシャル対応
	body := `{"frontend": 1,"backend": 2,"database": 3,"infrastructure": 3,"network": 2,"facilitation": 4}`
	fmt.Fprint(w, body)
}
