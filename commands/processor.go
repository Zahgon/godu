package commands

import (
	"sync"

	"github.com/viktomas/godu/files"
)

// ProcessFolder removes small files and sorts folder content based on accumulated size
func ProcessFolder(folder *files.File, limit int64) error { _ = "STUB: not implemented"; return nil }

// StartProcessing reads user commands and applies them to state
func StartProcessing(
	folder *files.File,
	commands <-chan Executer,
	states chan<- State,
	lastStateChan chan<- *State,
	wg *sync.WaitGroup,
) {
	_ = "STUB: not implemented"
	return
}
