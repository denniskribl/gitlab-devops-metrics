package main

import (
	"github.com/xanzy/go-gitlab"
	"log"
)

func getAllProjectsForTeam(team string) (m map[string]int) {

	listOptions := gitlab.ListOptions{PerPage: 100, Page: 1}

	p := &gitlab.ListGroupProjectsOptions{
		ListOptions:      listOptions,
		IncludeSubgroups: gitlab.Bool(true),
	}

	m = make(map[string]int)

	for {
		projects, resp, err := c.Groups.ListGroupProjects(team, p)

		if err != nil {
			log.Fatal(err)
		}

		for _, project := range projects {
			m[project.Name] = project.ID
		}

		if resp.CurrentPage >= resp.TotalPages {
			break
		}

		p.Page = resp.NextPage
	}

	return
}
