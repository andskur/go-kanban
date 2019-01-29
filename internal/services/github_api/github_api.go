package api

import (
	"context"

	"github.com/google/go-github/v21/github"
	"golang.org/x/oauth2"
)

// GitHub implements github Api client
type GitHub struct {
	*github.Client
}

// Authenticate create authorized GitHub client
func (gh *GitHub) Authenticate(ctx context.Context, token string) {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	gh.Client = github.NewClient(tc)
}

// FetchMilestones fetch milestones for given owner and repository
func (gh *GitHub) FetchMilestones(ctx context.Context, owner, repo string) ([]*github.Milestone, error) {
	milestones, _, err := gh.Client.Issues.ListMilestones(ctx, owner, repo, &github.MilestoneListOptions{})
	if err != nil {
		return nil, err
	}
	return milestones, nil
}

// FetchIssues fetch and filter issues for given milestone
func (gh *GitHub) FetchIssues(ctx context.Context, owner, repo, ms string) ([]*github.Issue, error) {
	options := &github.IssueListByRepoOptions{
		Milestone: ms,
		State:     "all",
	}
	issues, _, err := gh.Client.Issues.ListByRepo(ctx, owner, repo, options)
	if err != nil {
		return nil, err
	}
	return issues, nil
}
