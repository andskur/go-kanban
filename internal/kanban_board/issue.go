package kanban

import (
	"fmt"
	"time"

	"github.com/google/go-github/v21/github"
)

// Milestone represent structure of Issue
type Issue struct {
	Id     int64
	State  string
	Title  string
	Url    string
	Closed *time.Time
	Assignee
}

// Milestone represent structure of Issue Assignee user
type Assignee struct {
	AvatarUrl string
}

// NewIssue create new Issue structure
func NewIssue(rawIssue *github.Issue) *Issue {
	issue := &Issue{
		Id:     *rawIssue.ID,
		Title:  *rawIssue.Title,
		Url:    *rawIssue.URL,
		Closed: rawIssue.ClosedAt,
	}
	issue.SetState(*rawIssue.State, rawIssue.Assignee)
	issue.SetAssignee(rawIssue.User)
	return issue
}

// SetState set state for current issue
func (ii *Issue) SetState(state string, assignee *github.User) {
	switch {
	case state == "closed":
		ii.State = "completed"
	case assignee != nil:
		ii.State = "active"
	default:
		ii.State = "queued"
	}
}

// SetAssignee set assignee user for current issue
func (ii *Issue) SetAssignee(user *github.User) {
	avatar := fmt.Sprintf("%s?s=16", *user.AvatarURL)
	assignee := Assignee{avatar}
	ii.Assignee = assignee
}