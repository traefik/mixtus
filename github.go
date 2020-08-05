package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

type message struct {
	Title string
	Body  string
}

func createPullRequest(ctx context.Context, cfg Config, branchName string) error {
	client := newGitHubClient(ctx, cfg.Token)

	msg := message{
		Title: fmt.Sprintf("Update documentation for %s", cfg.Source.RepoName),
		Body: fmt.Sprintf(`Update documentation for [%s](https://github.com/%s/%s).`,
			cfg.Source.RepoName, cfg.Source.Owner, cfg.Source.RepoName),
	}

	newPR := &github.NewPullRequest{
		Title:               github.String(msg.Title),
		Head:                github.String(branchName),
		Base:                github.String("master"),
		Body:                github.String(msg.Body),
		MaintainerCanModify: github.Bool(true),
	}

	pullRequest, _, err := client.PullRequests.Create(ctx, cfg.Target.Owner, cfg.Target.RepoName, newPR)
	if err != nil {
		return err
	}

	log.Println(pullRequest.GetHTMLURL())

	labels := []string{"status/3-needs-merge", fmt.Sprintf("area/%s", cfg.Source.RepoName)}

	_, _, err = client.Issues.AddLabelsToIssue(ctx, cfg.Target.Owner, cfg.Target.RepoName, pullRequest.GetNumber(), labels)
	if err != nil {
		return err
	}

	return nil
}

func newGitHubClient(ctx context.Context, token string) *github.Client {
	if len(token) == 0 {
		return github.NewClient(nil)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	return github.NewClient(oauth2.NewClient(ctx, ts))
}
