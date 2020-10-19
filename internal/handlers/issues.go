package handlers

import (
	"context"
	mastermind "github.com/orlade/mastermind/gen/pkg/servers/Mastermind"
	"github.com/orlade/mastermind/gen/pkg/servers/Mastermind/github"
)

func GetReposIssuesList(ctx context.Context, req *mastermind.GetReposIssuesListRequest, client mastermind.GetReposIssuesListClient) (*[]mastermind.Issue, error) {
	downReq := &github.GetReposIssuesListRequest{
		Owner: req.Owner,
		Repo:  req.Repo,
	}
	downRes, err := client.GithubGetReposIssuesList(ctx, downReq)
	if err != nil {
		return nil, err
	}

	res := []mastermind.Issue{}
	for _, i := range *downRes {
		res = append(res, mastermind.Issue{
			ID:    i.ID,
			Title: i.Title,
		})
	}
	return &res, nil
}
