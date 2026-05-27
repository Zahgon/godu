package commands

import (
	"github.com/viktomas/godu/files"
)

// State represents system configuration after processing user input
type State struct {
	Folder      *files.File
	Selected    int
	history     map[*files.File]int // last cursor location in each folder
	MarkedFiles map[*files.File]struct{}
}

// Executer represents a user action triggered on a State
type Executer interface {
	Execute(State) (State, error)
}

// Enter is an action opening selected directory
type Enter struct{}

// GoBack is an action returning to parent directory
type GoBack struct{}

// Down is an action selecting next file in the list
type Down struct{}

// Up is an action selecting previous file in the list
type Up struct{}

// Mark is an action that saves current directory for later use
type Mark struct{}

func copyState(state State) State { _ = "STUB: not implemented"; return *new(State) }

func (d Down) Execute(oldState State) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

func (u Up) Execute(oldState State) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

func (e Enter) Execute(oldState State) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

func (GoBack) Execute(oldState State) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}

func (m Mark) Execute(oldState State) (State, error) {
	_ = "STUB: not implemented"
	return *new(State), nil
}
