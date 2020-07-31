package main

import (
	"bufio"
	"fmt"
	"github.com/xanzy/go-gitlab"
	"log"
	"os"
	"strings"
)

var (
	c             *gitlab.Client
	gitlabDomain  string
	gitlabBaseURL string
)

func init() {
	var err error
	const defaultGitlabBaseDomain string = "https://gitlab.com"

	gitlabDomain = os.Getenv("GITLAB_BASE_URL")
	if gitlabDomain == "" {
		gitlabDomain = defaultGitlabBaseDomain
	}

	gitlabBaseURL = gitlabDomain + "/api/v4"

	token := os.Getenv("GITLAB_TOKEN")
	if token == "" {
		fmt.Println("WARNING: the environment variable GITLAB_TOKEN is not set. only crawling public projects...")
	}

	c, err = gitlab.NewClient(token, gitlab.WithBaseURL(gitlabBaseURL))
	if err != nil {
		log.Fatalf("Failed to create GitLab API client: %v", err)
	}
}

func main() {
	//Read team
	r := bufio.NewReader(os.Stdin)
	fmt.Print("GitLab Team name: ")
	t, _ := r.ReadString('\n')
	team := strings.Trim(t, " \n")

	//Get all projects for team
	projects := getAllProjectsForTeam(team)

	commits := make(map[string]int)
	successfulPipelines := make(map[string]int)
	failedPipelines := make(map[string]int)
	pipelines := make(map[string]int)

	totalCommits := 0
	totalSuccessfulPipelines := 0
	totalFailedPipelines := 0
	totalPipelines := 0

	fmt.Println(fmt.Sprintf("total projects: %v", len(projects)))
	since, until := getLastMonth()
	fmt.Println(fmt.Sprintf("getting all metrics between %v and %v.... this may take a while", since, until))

	// todo: make concurrent

	for project, count := range projects {
		commits[project] = getCommitCountForLastMonth(count)
		totalCommits += commits[project]

		successfulPipelines[project] = getPipelineCountForLastMonth(count, gitlab.BuildState(gitlab.Success))
		totalSuccessfulPipelines += successfulPipelines[project]

		failedPipelines[project] = getPipelineCountForLastMonth(count, gitlab.BuildState(gitlab.Failed))
		totalFailedPipelines += failedPipelines[project]

		pipelines[project] = getPipelineCountForLastMonth(count, nil)
		totalPipelines += pipelines[project]
	}

	sortedPrintByType(commits, "commits")
	fmt.Println(fmt.Sprintf("total commits: %v", totalCommits))

	sortedPrintByType(successfulPipelines, "successful pipelines")
	fmt.Println(fmt.Sprintf("total successful pipelines: %v", totalSuccessfulPipelines))

	sortedPrintByType(failedPipelines, "failed pipelines")
	fmt.Println(fmt.Sprintf("total failed pipelines: %v", totalFailedPipelines))

	fmt.Println(fmt.Sprintf("total pipelines : %v", totalPipelines))
	fmt.Println(fmt.Sprintf("total change failure rate: %v %%", totalFailedPipelines*100/totalPipelines))

}
