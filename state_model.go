package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/viktomas/godu/commands"
	"github.com/viktomas/godu/interactive"
)

type visualState struct {
	folders        []interactive.Line
	selected       int
	xbound, ybound int
	screenHeight   int
}

func newVisualState(state commands.State, screenHeight int) visualState {
	_ = "STUB: not implemented"
	return *new(visualState)
}

func (vs visualState) GetCell(x, y int) (rune, tcell.Style, []rune, int) {
	_ = "STUB: not implemented"
	return 0,

		// return empty cell if we are asking for a line that doesn't exist
		*new(tcell.Style), nil, 0
}

// For some reason tcell is asking for cells below the viewport, we will return empty cell

// shifting the index enables displaying selected folders that would be otherwise hidden bellow the screen

func (vs visualState) GetBounds() (int, int) { _ = "STUB: not implemented"; return 0, 0 }

func (visualState) SetCursor(int, int) { _ = "STUB: not implemented"; return }

func (visualState) GetCursor() (int, int, bool, bool) {
	_ = "STUB: not implemented"
	return 0, 0, false, false
}

func (visualState) MoveCursor(offx, offy int) { _ = "STUB: not implemented"; return }
