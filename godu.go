package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/gdamore/tcell/v2"
	"github.com/viktomas/godu/commands"
	"github.com/viktomas/godu/files"
)

// the correct version is injected by `go build` command in release.sh script
var goduVersion = "master"

func main() {
	limit := flag.Int64("l", 10, "show only files larger than limit (in MB)")
	nullTerminate := flag.Bool("print0", false, "print null-terminated strings")
	version := flag.Bool("v", false, "show version")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: godu [OPTION]... [DIRECTORY]\nShow disk usage under DIRECTORY (. by default) interactively.\n\nOptions:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nThe currently selected file/folder can be marked/unmarked with the space key. Upon exiting, godu prints all marked files/folders to stdout. You can further process them with commands like xargs.\n\nFor example:\n\n# Show information of selected files\ngodu -print0 | xargs -0 ls -l\n\n# Delete selected files\ngodu -print0 | xargs -0 rm -rf\n\n# Move selected files to 'tmp' directory\ngodu -print0 | xargs -0 -I _ mv _ tmp\n")
	}
	flag.Parse()
	if *version {
		fmt.Printf("godu %s\n", goduVersion)
		os.Exit(0)
	}
	args := flag.Args()
	rootFolderName := "."
	if len(args) > 0 {
		rootFolderName = args[0]
	}
	rootFolderName, err := filepath.Abs(rootFolderName)
	if err != nil {
		log.Fatalln(err.Error())
	}
	progress := make(chan int)
	go reportProgress(progress)
	rootFolder := files.WalkFolder(rootFolderName, ioutil.ReadDir, ignoreBasedOnIgnoreFile(readIgnoreFile()), progress)
	rootFolder.Name = rootFolderName
	err = commands.ProcessFolder(rootFolder, *limit*files.MEGABYTE)
	if err != nil {
		log.Fatalln(err.Error())
	}
	s := initScreen()
	commandsChan := make(chan commands.Executer)
	states := make(chan commands.State)
	lastStateChan := make(chan *commands.State, 1)
	var wg sync.WaitGroup
	wg.Add(3)
	go commands.StartProcessing(rootFolder, commandsChan, states, lastStateChan, &wg)
	go interactiveFolder(s, states, &wg)
	go parseCommand(s, commandsChan, &wg)
	wg.Wait()
	s.Fini()
	lastState := <-lastStateChan
	printMarkedFiles(lastState, *nullTerminate)
}

func reportProgress(progress <-chan int) { _ = "STUB: not implemented"; return }

func printMarkedFiles(lastState *commands.State, nullTerminate bool) {
	_ = "STUB: not implemented"
	return
}

func initScreen() tcell.Screen { _ = "STUB: not implemented"; return *new(tcell.Screen) }
