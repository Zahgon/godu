package interactive

import (
	"github.com/viktomas/godu/files"
)

type byLength []string

func (l byLength) Len() int           { _ = "STUB: not implemented"; return 0 }
func (l byLength) Swap(i, j int)      { _ = "STUB: not implemented"; return }
func (l byLength) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// FilesAsSlice takes files from the map and returns a sorted slice of file paths.
func FilesAsSlice(in map[*files.File]struct{}) []string { _ = "STUB: not implemented"; return nil }

// sorting length of the path (assuming that we want to delete files in subdirs first)
// alphabetical sorting added for determinism (map keys doesn't guarantee order)
