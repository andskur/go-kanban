package kanban

import (
	"context"
	"github.com/andskur/kanban-board/config"
	"github.com/andskur/kanban-board/internal/services/github_api"
)

//TODO Get rid of that shit
//GitHub client vars
var (
	token  string
	client api.GitHub
)

// Board represent Kanban board data structure
type Board struct {
	Owner        string
	Repositories []string
	Milestones   chan []*Milestone
}

// NewBoard create new Board structure
func NewBoard(config *config.Config) (*Board, error) {
	token = config.AccessToken

	board := &Board{
		Owner:        config.Account,
		Repositories: config.Repositories,
	}

	err := board.GetMilestones()
	if err != nil {
		return nil, err
	}

	return board, nil
}

// GetMilestones get milestones for  KanbanBoard instance
func (board *Board) GetMilestones() error {
	ctx := context.Background()
	client.Authenticate(ctx, token)

	c := make(chan []*Milestone)
	for _, repo := range board.Repositories {
		go board.FetchData(c, ctx, repo)

		/*milestones, err := client.FetchMilestones(ctx, board.Owner, repo)
		if err != nil {
			return err
		}

		for _, v := range milestones {
			ms, err := NewMilestone(&ctx, v, board.Owner, repo)
			if err != nil {
				return err
			}
			board.Milestones = append(board.Milestones, ms)
		}*/
		board.Milestones = c
	}
	/*sort.Slice(board.Milestones, func(i, j int) bool {
		return board.Milestones[i].Title < board.Milestones[j].Title
	})*/
	return nil
}

func (board *Board) FetchData(data chan []*Milestone, ctx context.Context, repo string) error {
	var mss []*Milestone
	milestones, err := client.FetchMilestones(ctx, board.Owner, repo)
	if err != nil {
		return err
	}
	for _, v := range milestones {
		ms, err := NewMilestone(&ctx, v, board.Owner, repo)
		if err != nil {
			return err
		}
		mss = append(mss, ms)
	}

	data <- mss

	return nil
}
