package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type Note struct {
	ReasonForExam string
}

func main() {
	indexTemplatePage := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexTemplatePage.Execute(w, nil)
	})

	noteTemplate := template.Must(template.ParseFiles("note.html"))
	http.HandleFunc("/generate-note", func(w http.ResponseWriter, r *http.Request) {
		var note Note
		_ = json.NewDecoder(r.Body).Decode(&note)
		fmt.Println("Connected is", note)
		noteTemplate.Execute(w, note)
	})

	http.ListenAndServe(":8080", nil)
}
