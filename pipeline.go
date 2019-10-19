package main

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

// GetPieplineData from projects
func GetPieplineData(c *gitlab.Client, pid string) {
	pipelines, err := GetPipelines(c, pid)
	if err == nil {
		for _, pipeline := range pipelines {
			fmt.Println(pipeline.Ref)
			fmt.Println(pipeline.Status)
			fmt.Println(pipeline.ID)

			jobs, err := GetJobs(c, pid, pipeline.ID)
			if err == nil {
				for _, job := range jobs {
					fmt.Println(job.Name)
					fmt.Println(job.Status)
					fmt.Println(job.StartedAt)
					fmt.Println(job.User.Name)
					fmt.Println(job.User.AvatarURL)
				}
			}
		}
	}
}
