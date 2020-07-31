package main

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
	"log"
	"sort"
)

func getCommitCountForLastMonth(projectId int) (commitCount int) {
	since, until := getLastMonth()
	listOptions := gitlab.ListOptions{PerPage: 100, Page: 1}

	options := &gitlab.ListCommitsOptions{
		ListOptions: listOptions,
		Since:       &since,
		Until:       &until,
		All:         gitlab.Bool(true),
		WithStats:   gitlab.Bool(false),
	}

	commitCount = 0

	for {
		commits, resp, err := c.Commits.ListCommits(projectId, options)

		if err != nil {
			log.Fatal(err)
		}
		commitCount += len(commits)

		if resp.CurrentPage >= resp.TotalPages {
			break
		}
		options.Page = resp.NextPage
	}

	return
}

func sortedPrintByType(m map[string]int, metricType string) {
	n := map[int][]string{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	for _, k := range a {
		for _, s := range n[k] {
			fmt.Printf("%v of %s last month: %d\n", metricType, s, k)
		}
	}
}
