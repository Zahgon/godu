package main

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/viktomas/godu/commands"
)

func interactiveFolder(s tcell.Screen, states chan commands.State, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

func printOptions(state commands.State, s tcell.Screen) { _ = "STUB: not implemented"; return }

// Subtract a row from screen height for the status bar
