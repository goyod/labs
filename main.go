package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/goyod/labs/fizzbuzz"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz/{number}", fizzbuzzHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	n, err := strconv.Atoi(vars["number"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, fizzbuzz.FizzBuzz(n))
}
