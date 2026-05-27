package main

import (
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/viktomas/godu/commands"
)

func parseCommand(s tcell.Screen, commandsChan chan commands.Executer, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}
