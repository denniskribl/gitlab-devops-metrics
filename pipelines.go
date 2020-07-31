package main

import (
	"github.com/xanzy/go-gitlab"
)

func getPipelineCountForLastMonth(projectId int, buildState *gitlab.BuildStateValue) (count int) {
	since, until := getLastMonth()
	listOptions := gitlab.ListOptions{PerPage: 100, Page: 1}

	options := &gitlab.ListProjectPipelinesOptions{
		ListOptions:   listOptions,
		UpdatedAfter:  &since,
		UpdatedBefore: &until,
	}

	if buildState != nil {
		options.Status = buildState
	}

	count = 0

	for {
		pipelines, resp, _ := c.Pipelines.ListProjectPipelines(projectId, options)
		count += len(pipelines)

		if resp.CurrentPage >= resp.TotalPages {
			break
		}
		options.Page = resp.NextPage
	}
	return
}
