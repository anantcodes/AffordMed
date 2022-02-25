package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"net/http"

	"github.com/gorilla/mux"
)

type Result struct {
	keyword string
	status  string
	prefix  string
}

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Server is up and running...")
	})

	router.HandleFunc("/prefixes", Prefixes).Methods("GET")

	log.Println("Server Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func Prefixes(rw http.ResponseWriter, r *http.Request) {

	keys := [5]string{"bonfire", "cardio", "case", "character", "bonsai"}

	rw.Header().Set("Content-type", "application/json")
	keywords := r.URL.Query().Get("keywords")

	if r.URL == nil {
		json.NewEncoder(rw).Encode("Please send some data")
		return
	}

	res := []*Result{}

	s := strings.Split(keywords, ",")
	for _, s := range s {
		for _, v := range keys {
			if v == s {
				res = append(res, &Result{keyword: s, status: "", prefix: ""})
			}
		}
	}

	j, _ := json.Marshal(&res)
	fmt.Println(string(j))
	json.NewEncoder(rw).Encode(string(j))
}