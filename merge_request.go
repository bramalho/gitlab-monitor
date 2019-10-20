package main

import (
	"time"

	"github.com/xanzy/go-gitlab"
)

// MergeRequest struct
type MergeRequest struct {
	Project      Project
	Title        string
	WIP          bool
	CreatedAt    *time.Time
	AuthorName   string
	AuthorAvatar string
	Upvotes      int
	Downvotes    int
	Color        string
}

// GetMergeRequestData from projects
func GetMergeRequestData(c *gitlab.Client, p Project) []MergeRequest {
	mergeRequests := []MergeRequest{}

	mrs, err := GetMergeRequests(c, p.ID)
	if err == nil {
		for _, mr := range mrs {
			color := "secondary"

			if mr.Upvotes > 0 {
				color = "success"
			}
			if mr.Downvotes > mr.Upvotes {
				color = "danger"
			}
			if mr.WorkInProgress {
				color = "warning"
			}

			mergeRequests = append(mergeRequests, MergeRequest{
				Project:      p,
				Title:        mr.Title,
				WIP:          mr.WorkInProgress,
				CreatedAt:    mr.CreatedAt,
				AuthorName:   mr.Author.Name,
				AuthorAvatar: mr.Author.AvatarURL,
				Upvotes:      mr.Upvotes,
				Downvotes:    mr.Downvotes,
				Color:        color,
			})
		}
	}

	return mergeRequests
}
