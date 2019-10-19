package main

import (
	"fmt"
	"log"

	"github.com/bramalho/gitlab-monitor/client"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	fmt.Println("It works!")

	pid := "1945"

	git := client.GetClient()

	project, err := client.GetProject(git, pid)
	if err == nil {
		fmt.Println(project.NameWithNamespace)
		fmt.Println(project.AvatarURL)
	}

	mergeRequests, err := client.GetMergeRequests(git, pid)
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

	pipelines, err := client.GetPipelines(git, pid)
	if err == nil {
		for _, pipeline := range pipelines {
			fmt.Println(pipeline.Ref)
			fmt.Println(pipeline.Status)
			fmt.Println(pipeline.ID)

			jobs, err := client.GetJobs(git, pid, pipeline.ID)
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
