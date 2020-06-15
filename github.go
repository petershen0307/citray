package main

import (
	"context"
	"log"
	"path"
	"strings"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

type prStruct struct {
	url string
}

type githubClientWrapper struct {
	client *github.Client
	ctx    context.Context
}

func (c *githubClientWrapper) init(ctx context.Context, token, enterpriceURL string) {
	c.ctx = ctx
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(c.ctx, ts)
	var err error
	if enterpriceURL == "" {
		c.client = github.NewClient(tc)
	} else {
		c.client, err = github.NewEnterpriseClient(enterpriceURL, enterpriceURL, tc)
		if err != nil {
			log.Println("NewEnterpriseClient() error:", err)
		}
	}
}

func (c *githubClientWrapper) queryPRFromSetting(setting Setting) map[string][]prStruct {
	repositories := make(map[string][]prStruct)
	for _, repo := range setting.Github.Repositories {
		prList := c.getSpecificReviewerPR(repo.Owner, repo.RepoName, setting.Github.UserName)
		repositories[path.Join(repo.Owner, repo.RepoName)] = append(repositories[path.Join(repo.Owner, repo.RepoName)], prList...)
	}
	return repositories
}

func (c *githubClientWrapper) getSpecificReviewerPR(owner, repoName, prReviewer string) []prStruct {
	opt := &github.PullRequestListOptions{State: "open", Sort: "created", Direction: "desc"}
	// currently not support pagination
	pullRequests, response, err := c.client.PullRequests.List(c.ctx, owner, repoName, opt)
	prList := []prStruct{}
	if err != nil {
		log.Println(err)
	}
	if response != nil {
		log.Println(response.Status)
	}
	log.Println("pr size:", len(pullRequests))
	for _, pr := range pullRequests {
		log.Println("repo:", path.Join(owner, repoName), "pr url:", pr.GetURL())
		if len(pr.RequestedReviewers) == 0 {
			prList = append(prList, prStruct{
				url: pr.GetURL(),
			})
		}
		for _, reviewer := range pr.RequestedReviewers {
			if strings.ToLower(reviewer.GetLogin()) == strings.ToLower(prReviewer) {
				prList = append(prList, prStruct{
					url: pr.GetURL(),
				})
			}
			log.Println("repo:", path.Join(owner, repoName), "reviewer:", reviewer.GetLogin())
		}
	}
	return prList
}
