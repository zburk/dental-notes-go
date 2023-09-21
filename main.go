package main

import (
	"net/http"
	"text/template"
)

type Note struct {
	ReasonForExam           string
	NextVisit               string
	Xrays                   string
	PerioDx                 string
	ClinicalCaries          string
	ToothNumbers            string
	ToothNumbersAndSurfaces string
}

func main() {
	indexTemplatePage := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexTemplatePage.Execute(w, nil)
	})

	noteSubFieldsTemplate := template.Must(template.ParseFiles("templates/noteSubFields.html"))
	http.HandleFunc("/generate-subfields", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			ReasonForExam string
		}{ReasonForExam: r.PostFormValue("ReasonForVisit")}

		noteSubFieldsTemplate.Execute(w, data)
	})

	noteTemplate := template.Must(template.ParseFiles("templates/note.html"))
	http.HandleFunc("/generate-note", func(w http.ResponseWriter, r *http.Request) {
		note := Note{
			ReasonForExam:           r.PostFormValue("ReasonForVisit"),
			NextVisit:               r.PostFormValue("NextVisit"),
			Xrays:                   r.PostFormValue("Xrays"),
			PerioDx:                 r.PostFormValue("PerioDx"),
			ClinicalCaries:          r.PostFormValue("ClinicalCaries"),
			ToothNumbers:            r.PostFormValue("ToothNumbers"),
			ToothNumbersAndSurfaces: r.PostFormValue("ToothNumbersAndSurfaces"),
		}

		noteTemplate.Execute(w, note)
	})

	http.ListenAndServe(":8080", nil)
}
