package kanban

import (
	"context"
	"math"
	"strconv"

	"github.com/google/go-github/v21/github"
)

// Milestone represent structure of repository milestone data
type Milestone struct {
	Id    int
	Title string
	Url   string
	*Progress
	Issues
}

// Progress represent progress of given milestone
type Progress struct {
	Total     int
	Complete  int
	Remaining int
	Percent   int
}

// Progress represent issues of given milestone
type Issues struct {
	Queued    []*Issue
	Active    []*Issue
	Completed []*Issue
}

// NewMilestone create new Milestone structure
func NewMilestone(ctx *context.Context, rawMs *github.Milestone, owner, repo string, pausedLabels []string) (*Milestone, error) {
	ms := &Milestone{
		Id:    *rawMs.Number,
		Title: *rawMs.Title,
		Url:   *rawMs.HTMLURL,
	}

	ms.GetProgress(float64(*rawMs.ClosedIssues), float64(*rawMs.OpenIssues))

	err := ms.GetIssues(ctx, owner, repo, pausedLabels)
	if err != nil {
		return nil, err
	}

	return ms, nil
}

// GetMilestones get issues for Milestone instance
func (ms *Milestone) GetIssues(ctx *context.Context, owner, repo string, pausedLabels []string) error {
	issues, err := client.FetchIssues(*ctx, owner, repo, strconv.Itoa(ms.Id))
	if err != nil {
		return err
	}

	for _, ii := range issues {
		issue := NewIssue(ii, pausedLabels)
		switch issue.State {
		case "queued":
			ms.Issues.Queued = append(ms.Issues.Queued, issue)
		case "active":
			ms.Issues.Active = append(ms.Issues.Active, issue)
		case "completed":
			ms.Issues.Completed = append(ms.Issues.Completed, issue)
		}
	}

	return nil
}

// GetProgress calculate current Milestone issues progress
func (ms *Milestone) GetProgress(complete, remaining float64) {
	total := complete + remaining
	percent := 0
	if total > 0 {
		percent = int(math.Round(complete / total * 100))
	}

	prog := &Progress{int(total), int(complete), int(remaining), percent}
	ms.Progress = prog
}
