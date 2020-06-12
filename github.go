package main

import (
	"context"
	"log"

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
	c.client, err = github.NewEnterpriseClient(enterpriceURL, enterpriceURL, tc)
	if err != nil {
		log.Println("NewEnterpriseClient() error:", err)
	}
}

func (c *githubClientWrapper) getSpecificReviewerPR(owner, repoName, prReviewer string) []prStruct {
	opt := &github.PullRequestListOptions{State: "open", Sort: "created", Direction: "desc"}
	// currently not support pagination
	pullRequests, response, err := c.client.PullRequests.List(c.ctx, owner, repoName, opt)
	prList := []prStruct{}
	log.Println(err)
	if response != nil {
		log.Println(response.Status)
	}
	log.Println(len(pullRequests))
	for _, pr := range pullRequests {
		log.Println(pr.GetURL())
		if len(pr.RequestedReviewers) == 0 {
			prList = append(prList, prStruct{
				url: pr.GetURL(),
			})
		}
		for _, reviewer := range pr.RequestedReviewers {
			if reviewer.GetLogin() == prReviewer {
				prList = append(prList, prStruct{
					url: pr.GetURL(),
				})
			}
			log.Println(reviewer.GetLogin())
		}
	}
	return prList
}
