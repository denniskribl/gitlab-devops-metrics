# gitlab-devops-metrics

## About

This project should help teams get some insights about their DevOps metrics.  
It uses GitLab and its API to get those metrics. Currently only crawls GitLab for metrics of the last month

## Usage

compile the code e.g. on linux:

```bash
GOOS=linux GOARCH=amd64 go build -o builds/get-devops-metrics
```

execute the binary and type in your group/team name:

```bash
./get-devops-metrics 
```

## Providing an Auth Token

you can set the environment variable `GITLAB_TOKEN` to also access non public repositories e.g.

```bash
GITLAB_TOKEN=my-super-secret-token ./get-devops-metrics
```

## Custom GitLab instance

if you want to run the crawler against a on-prem gitlab instance you can provide the environment variable 
`GITLAB_BASE_URL` an example could look like this:

```bash
GITLAB_BASE_URL="https://my-gitlabs-cooler-than-yours.mycompany.com" ./get-devops-metrics
```

## Roadmap

- [x] commits
- [x] deployments
- [ ] set custom timespan
- [ ] make crawling parallel
- [ ] better output formatting
