package kanban

import (
	"context"
	"fmt"
	"log"
	"sort"
	"sync"

	"github.com/google/go-github/v21/github"

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
	Milestones   []*Milestone
}

// NewBoard create new Board structure
func NewBoard(config *config.Config) (*Board, error) {
	token = config.AccessToken

	board := &Board{
		Owner:        config.Account,
		Repositories: config.Repositories,
	}

	err := board.CreateBoard()
	if err != nil {
		return nil, err
	}

	return board, nil
}

// GetMilestones async get milestones for KanbanBoard instance
func (board *Board) CreateBoard() error {
	ctx := context.Background()
	client.Authenticate(ctx, token)

	var wg sync.WaitGroup

	// first, fetch milestones for each repository
	for _, repo := range board.Repositories {
		milestones, err := client.FetchMilestones(ctx, board.Owner, repo)
		if err != nil {
			return err
		}

		// next, format each milestone and fetch its issues
		for _, v := range milestones {
			wg.Add(1)
			go board.PrepareMilestones(&wg, ctx, v, repo)
		}

		// waiting data from all milestones
		wg.Wait()
	}
	return nil
}

// PrepareMilestones async format given milestones and fetch its issues
func (board *Board) PrepareMilestones(wg *sync.WaitGroup, ctx context.Context, rawmS *github.Milestone, repo string) {
	defer wg.Done()

	fmt.Printf("Start fetching for %s milestone\n", rawmS.GetTitle())
	milestone, err := NewMilestone(&ctx, rawmS, board.Owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	board.Milestones = append(board.Milestones, milestone)
	issuesCount := len(milestone.Issues.Queued) + len(milestone.Issues.Completed) + len(milestone.Issues.Active)
	fmt.Printf("Finish fetching %d issues for %s milestone\n", issuesCount, milestone.Title)
}

// SortMilestones sort board milestones in alphabet order by its Title
func (board *Board) SortMilestones() {
	sort.Slice(board.Milestones, func(i, j int) bool {
		return board.Milestones[i].Title < board.Milestones[j].Title
	})
}
