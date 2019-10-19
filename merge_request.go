package main

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

// GetMergeRequestData from projects
func GetMergeRequestData(c *gitlab.Client, pid string) {
	mergeRequests, err := GetMergeRequests(c, pid)
	if err == nil {
		for _, mergeRequest := range mergeRequests {
			fmt.Println(mergeRequest.Title)
			fmt.Println(mergeRequest.WorkInProgress)
			fmt.Println(mergeRequest.CreatedAt)
			fmt.Println(mergeRequest.Author.Name)
			fmt.Println(mergeRequest.Author.AvatarURL)
			fmt.Println(mergeRequest.Upvotes)
			fmt.Println(mergeRequest.Downvotes)
		}
	}
}
