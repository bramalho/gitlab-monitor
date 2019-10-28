package main

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/bradfitz/slice"
)

// HomeHandler controller method
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	git := GetClient()

	projects := GetProjectData(git)

	mergeRequests := []MergeRequest{}
	pipelines := []Pipeline{}

	var wg sync.WaitGroup
	wg.Add(len(projects))

	for _, project := range projects {
		go func(project Project) {
			defer wg.Done()
			mrs := GetMergeRequestData(git, project)
			for _, mr := range mrs {
				mergeRequests = append(mergeRequests, mr)
			}

			ps := GetPieplineData(git, project)
			for _, p := range ps {
				pipelines = append(pipelines, p)
			}
		}(project)
	}

	wg.Wait()

	slice.Sort(mergeRequests[:], func(i, j int) bool {
		return mergeRequests[i].CreatedAt.UnixNano() > mergeRequests[j].CreatedAt.UnixNano()
	})

	slice.Sort(pipelines[:], func(i, j int) bool {
		if pipelines[i].StartedAt != nil && pipelines[j].StartedAt != nil {
			return pipelines[i].StartedAt.UnixNano() > pipelines[j].StartedAt.UnixNano()
		}
		return false
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
