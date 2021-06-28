package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	ADDRESS = ":8080"
)

var (
	ACCEPTED_METHODS = []string{
		http.MethodPost,
		http.MethodGet,
		http.MethodPatch,
		http.MethodPut,
		http.MethodDelete,
	}
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Implement me!, I'm the %s method on %s path", r.Method, r.URL.Path)
}

func peopleHandler(people []person) http.HandlerFunc {
	C := defaultHandler
	R := func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(people); err != nil {
			http.Error(w, jsonError("can't encode json"), http.StatusInternalServerError)
		}
	}
	U := defaultHandler
	D := defaultHandler

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Default behavior, can be updated in the handlers of the specific methods
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		switch r.Method {
		case http.MethodPost:
			C(w, r)
		case http.MethodGet:
			R(w, r)
		case http.MethodPatch:
			fallthrough
		case http.MethodPut:
			U(w, r)
		case http.MethodDelete:
			D(w, r)
		default:
			w.Header().Set("Allow", strings.Join(ACCEPTED_METHODS, ","))
			http.Error(w, "we only accept CRUD methods", http.StatusMethodNotAllowed)
		}
	})
}

func Start(ch chan<- error) {
	people, err := getPeople()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/people", peopleHandler(people))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/people", http.StatusMovedPermanently)
	})

	log.Println("HTTP server listening on " + ADDRESS)
	ch <- http.ListenAndServe(ADDRESS, nil)
}

func jsonError(s string) string {
	msg := map[string]string{
		"error": s,
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return "something went very wrong. :-("
	}

	return string(data)
}
