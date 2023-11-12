package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var port = ":8080"

// function init akan otomatis dipanggil ketika package tersebut diimport / di run
func init() {
	students = append(students, &Student{ID: "S001", Name: "Bourne", Grade: 1})
	students = append(students, &Student{ID: "S002", Name: "Ethan", Grade: 2})
	students = append(students, &Student{ID: "S003", Name: "Wick", Grade: 3})
	students = append(students, &Student{ID: "S004", Name: "fani", Grade: 4})
}

func main() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/student", ActionStudent)

	var handler http.Handler = mux
	handler = MiddlewareAuth(handler)
	handler = MiddlewareAllowOnlyGet(handler)

	server := new(http.Server)
	server.Addr = port
	server.Handler = handler

	log.Printf("server running on localhost%s\n", port)
	server.ListenAndServe()
}

func ActionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		result := SelectStudent(id)
		if result == nil {
			OutputJSON(w, map[string]string{"message": "student not found"})
			return
		}

		OutputJSON(w, SelectStudent(id))
	}

	OutputJSON(w, GetStudents())
}

func OutputJSON(w http.ResponseWriter, o any) {

	// json.Marshal => konversi ke json byte
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
