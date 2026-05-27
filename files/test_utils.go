package files

// NewTestFolder is providing easy interface to create folders for automated tests
// Never use in production code!
func NewTestFolder(name string, files ...*File) *File { _ = "STUB: not implemented"; return nil }

// NewTestFile provides easy interface to create files for automated tests
// Never use in production code!
func NewTestFile(name string, size int64) *File { _ = "STUB: not implemented"; return nil }

// FindTestFile helps testing by returning first occurrence of file with given name.
// Never use in production code!
func FindTestFile(folder *File, name string) *File { _ = "STUB: not implemented"; return nil }
