package main

import (
	"html/template"
	"net/http"

	"github.com/bradfitz/slice"
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

	slice.Sort(mergeRequests[:], func(i, j int) bool {
		return mergeRequests[i].CreatedAt.UnixNano() > mergeRequests[j].CreatedAt.UnixNano()
	})

	slice.Sort(pipelines[:], func(i, j int) bool {
		return pipelines[i].StartedAt.UnixNano() > pipelines[j].StartedAt.UnixNano()
	})

	tmpl := template.Must(template.ParseFiles("web/template/index.html"))
	tmpl.Execute(w, struct {
		MergeRequests []MergeRequest
		Pipelines     []Pipeline
	}{
		mergeRequests,
		pipelines,
	})
}
