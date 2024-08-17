package github

import (
	"encoding/json"
	"fmt"
)

type UserActivity struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Actor struct {
		Id           int    `json:"id"`
		DisplayLogin string `json:"display_login"`
	}
	Repo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	Payload struct {
		RepoId int    `json:"repository_id"`
		Action string `json:"action"`
		Issue  struct {
			Id int `json:"id"`
		} `json:"issue"`
		PULLRequest struct {
			Id int `json:"id"`
		} `json:"pull_request"`
		Commits []struct {
			Sha     string `json:"sha"`
			Message string `json:"message"`
			URL     string `json:"url"`
		} `json:"commits"`
	}
}

func PrintUserActivity(username string) {
	response := requestGithub("users/" + username + "/events")
	userActivity := []UserActivity{}
	json.Unmarshal(response, &userActivity)

	for _, activity := range userActivity {
		printActivity(activity)
		fmt.Println()
	}
}

func printActivity(activity UserActivity) {
	switch activity.Type {
	case "PushEvent":
		fmt.Printf("- Pushed %d commits to %s.", len(activity.Payload.Commits), activity.Repo.Name)
	case "CreateEvent":
		fmt.Printf("- Created %s.", activity.Repo.Name)
	case "DeleteEvent":
		fmt.Printf("- Deleted %s.", activity.Repo.Name)
	case "ForkEvent":
		fmt.Printf("- Forked %s.", activity.Repo.Name)
	case "IssuesEvent":
		fmt.Printf("- %s issue #%d on %s.", activity.Payload.Action, activity.Payload.Issue.Id, activity.Repo.Name)
	case "IssueCommentEvent":
		fmt.Printf("- Commented on issue #%d on %s.", activity.Payload.Issue.Id, activity.Repo.Name)
	case "WatchEvent":
		fmt.Printf("- Starred %s.", activity.Repo.Name)
	case "PullRequestEvent":
		fmt.Printf("- %s pull request #%d on %s.", activity.Payload.Action, activity.Payload.PULLRequest.Id, activity.Repo.Name)
	case "PullRequestReviewCommentEvent":
		fmt.Printf("- Commented on pull request #%d on %s.", activity.Payload.PULLRequest.Id, activity.Repo.Name)
	}
}
