package main

import (
	"os"
	"strings"

	"github.com/xanzy/go-gitlab"
)

// Project struct
type Project struct {
	ID     string
	Name   string
	Avatar string
}

func getProjetIds() []string {
	projects := []string{}
	if data, exists := os.LookupEnv("PROJECTS"); exists {
		result := strings.Split(data, ",")
		for i := range result {
			if len(result[i]) > 0 {
				projects = append(projects, result[i])
			}
		}
	}

	return projects
}

// GetProjectData from projects
func GetProjectData(c *gitlab.Client) []Project {
	projects := []Project{}

	for _, pid := range getProjetIds() {
		project, err := GetProject(c, pid)
		if err == nil {
			projects = append(projects, Project{
				ID:     pid,
				Name:   project.NameWithNamespace,
				Avatar: project.AvatarURL,
			})
		}
	}

	return projects
}
