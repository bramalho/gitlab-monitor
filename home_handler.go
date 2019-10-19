package main

import (
	"html/template"
	"net/http"
)

// HomeHandler controller method
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	pid := "1945"

	git := GetClient()

	GetProject(git, pid)
	GetMergeRequestData(git, pid)
	GetPipelines(git, pid)

	tmpl := template.Must(template.ParseFiles("web/template/index.html"))
	tmpl.Execute(w, nil)
}
