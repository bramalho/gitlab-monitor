package main

import (
	"time"

	"github.com/xanzy/go-gitlab"
)

// Pipeline struct
type Pipeline struct {
	Project    Project
	ID         int
	Ref        string
	StartedAt  *time.Time
	UserName   string
	UserAvatar string
	Status     string
	Color      string
	Jobs       []Job
}

// Job struct
type Job struct {
	Name   string
	Status string
	Color  string
	Icon   string
}

// GetPieplineData from projects
func GetPieplineData(c *gitlab.Client, proj Project) []Pipeline {
	pipelines := []Pipeline{}

	ps, err := GetPipelines(c, proj.ID)
	if err == nil {
		for _, p := range ps {
			var startedAt *time.Time = nil
			var userName string = ""
			var userAvatar string = ""

			jobs := []Job{}
			js, err := GetJobs(c, proj.ID, p.ID)
			if err == nil {
				for _, j := range js {
					if startedAt == nil {
						startedAt = j.StartedAt
					}
					if userName == "" {
						userName = j.User.Name
					}
					if userAvatar == "" {
						userAvatar = j.User.AvatarURL
					}

					color := "secondary"
					icon := "pause"

					if j.Status == "running" {
						color = "warning"
						icon = "play"
					}
					if j.Status == "success" {
						color = "success"
						icon = "check"
					}
					if j.Status == "failed" {
						color = "danger"
						icon = "times"
					}

					jobs = append(jobs, Job{
						Name:   j.Name,
						Status: j.Status,
						Color:  color,
						Icon:   icon,
					})
				}
			}

			color := "warning"

			if p.Status == "success" {
				color = "success"
			}
			if p.Status == "failed" {
				color = "danger"
			}
			if p.Status == "manual" {
				color = "info"
			}

			pipelines = append(pipelines, Pipeline{
				Project:    proj,
				ID:         p.ID,
				Ref:        p.Ref,
				StartedAt:  startedAt,
				UserName:   userName,
				UserAvatar: userAvatar,
				Status:     p.Status,
				Color:      color,
				Jobs:       jobs,
			})
		}
	}

	return pipelines
}
