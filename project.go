package main

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

// GetProjectData from projects
func GetProjectData(c *gitlab.Client, pid string) {
	project, err := GetProject(c, pid)
	if err == nil {
		fmt.Println(project.NameWithNamespace)
		fmt.Println(project.AvatarURL)
	}
}
