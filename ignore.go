package main

import (
	"github.com/viktomas/godu/files"
)

func readIgnoreFile() []string { _ = "STUB: not implemented"; return nil }

func ignoreBasedOnIgnoreFile(ignoreFile []string) files.ShouldIgnoreFolder {
	_ = "STUB: not implemented"
	return *new(files.ShouldIgnoreFolder)
}
