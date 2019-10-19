package client

import (
	"os"

	"github.com/xanzy/go-gitlab"
)

// GetClient for gitlab
func GetClient() *gitlab.Client {
	url, _ := os.LookupEnv("GITLAB_URL")
	token, _ := os.LookupEnv("GITLAB_TOKEN")

	git := gitlab.NewClient(nil, token)
	git.SetBaseURL(url)

	return git
}

// GetProject details
func GetProject(c *gitlab.Client, pid string) (*gitlab.Project, error) {
	options := &gitlab.GetProjectOptions{}

	project, _, err := c.Projects.GetProject(pid, options)

	return project, err
}

// GetMergeRequests for project
func GetMergeRequests(c *gitlab.Client, pid string) ([]*gitlab.MergeRequest, error) {
	options := &gitlab.ListProjectMergeRequestsOptions{State: gitlab.String("opened")}

	mergeRequests, _, err := c.MergeRequests.ListProjectMergeRequests(pid, options)

	return mergeRequests, err
}

// GetPipelines for project
func GetPipelines(c *gitlab.Client, pid string) (gitlab.PipelineList, error) {
	listOptions := gitlab.ListOptions{PerPage: 5}
	options := &gitlab.ListProjectPipelinesOptions{ListOptions: listOptions}

	pipelines, _, err := c.Pipelines.ListProjectPipelines(pid, options)

	return pipelines, err
}

// GetJobs for project
func GetJobs(c *gitlab.Client, pid string, pipleneID int) ([]*gitlab.Job, error) {
	options := &gitlab.ListJobsOptions{}

	jobs, _, err := c.Jobs.ListPipelineJobs(pid, pipleneID, options)

	return jobs, err
}
