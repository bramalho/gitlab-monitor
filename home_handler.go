package main

import (
	"html/template"
	"net/http"
)

// HomeHandler controller method
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	git := GetClient()

	projects := GetProjectData(git)
	mergeRequests := []MergeRequest{}
	pipelines := []Pipeline{}
	for _, project := range projects {
		mrs := GetMergeRequestData(git, project)
		for _, mr := range mrs {
			mergeRequests = append(mergeRequests, mr)
		}

		ps := GetPieplineData(git, project)
		for _, p := range ps {
			pipelines = append(pipelines, p)
		}
	}

	tmpl := template.Must(template.ParseFiles("web/template/index.html"))
	tmpl.Execute(w, struct {
		MergeRequests []MergeRequest
		Pipelines     []Pipeline
	}{
		mergeRequests,
		pipelines,
	})
}
