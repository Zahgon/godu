package interactive

import (
	"github.com/viktomas/godu/files"
)

// Line represents row of text in folder UI contains info about subfile
type Line struct {
	Text     []rune
	IsMarked bool
}

// Status contain info about size of all files in current godu instance
// and size of the files marked by user
type Status struct {
	Total    string
	Selected string
}

// ReportStatus reads through the folder structure and produces Status
func ReportStatus(file *files.File, markedFiles *map[*files.File]struct{}) Status {
	_ = "STUB: not implemented"
	return *new(Status)
}

func parentMarked(file *files.File, markedFiles *map[*files.File]struct{}) bool {
	_ = "STUB: not implemented"
	return false
}

// ReportFolder converts all subfiles into UI lines
func ReportFolder(folder *files.File, markedFiles map[*files.File]struct{}) []Line {
	_ = "STUB: not implemented"
	return nil
}

func formatBytes(bytesInt int64) string { _ = "STUB: not implemented"; return "" }
