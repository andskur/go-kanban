package kanban

import (
	"context"
	"log"
	"sort"
	"sync"
	"time"

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
	PausedLabels []string
}

var mutex sync.Mutex

// NewBoard create new Board structure
func NewBoard(config *config.Config) (*Board, error) {
	token = config.AccessToken

	board := &Board{
		Owner:        config.Account,
		Repositories: config.Repositories,
		PausedLabels: config.PausedLabels,
	}

	err := board.CreateBoard()
	if err != nil {
		return nil, err
	}

	board.SortMilestones()

	return board, nil
}

// GetMilestones async get milestones for KanbanBoard instance
func (board *Board) CreateBoard() error {
	ctx := context.Background()
	client.Authenticate(ctx, token)

	var wg sync.WaitGroup

	// first, fetch milestones for each repository
	for _, repo := range board.Repositories {
		wg.Add(1)
		go board.FetchMilestones(&wg, ctx, repo)
	}
	// waiting data from all repositories
	wg.Wait()
	return nil
}

// FetchMilestones fetch milestones data for given repository
func (board *Board) FetchMilestones(wg *sync.WaitGroup, ctx context.Context, repo string) {
	defer wg.Done()

	// first, fetch milestones for each repository
	milestones, err := client.FetchMilestones(ctx, board.Owner, repo)
	if err != nil {
		log.Println(err)
	}

	var wgSub sync.WaitGroup

	//next,  format each milestone and fetch its issues
	for _, v := range milestones {
		wgSub.Add(1)
		go board.PrepareMilestones(&wgSub, ctx, v, repo)
	}

	// waiting data from all milestones
	wgSub.Wait()
}

// PrepareMilestones async format given milestones and fetch its issues
func (board *Board) PrepareMilestones(wg *sync.WaitGroup, ctx context.Context, ms *github.Milestone, repo string) {
	defer wg.Done()

	start := time.Now()
	log.Printf("Start fetching issues for %s milestone\n", ms.GetTitle())
	milestone, err := NewMilestone(&ctx, ms, board.Owner, repo, board.PausedLabels)
	if err != nil {
		log.Println(err)
	}

	mutex.Lock()
	board.Milestones = append(board.Milestones, milestone)
	mutex.Unlock()

	issuesCount := len(milestone.Issues.Queued) + len(milestone.Issues.Completed) + len(milestone.Issues.Active)
	diff := time.Now().Sub(start)
	log.Printf("Finish fetching %d issues for %s mileston in %s \n", issuesCount, milestone.Title, diff)
}

// SortMilestones sort board milestones in alphabet order by its Title
func (board *Board) SortMilestones() {
	sort.Slice(board.Milestones, func(i, j int) bool {
		return board.Milestones[i].Title < board.Milestones[j].Title
	})
}
