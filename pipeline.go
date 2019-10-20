package main

import (
	"time"

	"github.com/xanzy/go-gitlab"
)

// Pipeline struct
type Pipeline struct {
	ID     int
	Ref    string
	Status string
	Jobs   []Job
}

// Job struct
type Job struct {
	Name       string
	Status     string
	StartedAt  *time.Time
	UserName   string
	UserAvatar string
}

// GetPieplineData from projects
func GetPieplineData(c *gitlab.Client, proj Project) []Pipeline {
	pipelines := []Pipeline{}

	ps, err := GetPipelines(c, proj.ID)
	if err == nil {
		for _, p := range ps {
			jobs := []Job{}
			js, err := GetJobs(c, proj.ID, p.ID)
			if err == nil {
				for _, j := range js {
					jobs = append(jobs, Job{
						Name:       j.Name,
						Status:     j.Status,
						StartedAt:  j.StartedAt,
						UserName:   j.User.Name,
						UserAvatar: j.User.AvatarURL,
					})
				}
			}

			pipelines = append(pipelines, Pipeline{
				ID:     p.ID,
				Ref:    p.Ref,
				Status: p.Status,
				Jobs:   jobs,
			})
		}
	}

	return pipelines
}
