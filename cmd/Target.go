package cmd

// A Target is an evaluation target, capabile of processing commands
type Target interface {
	// Load requests to load data files from two paths.
	Load(path1, path2 string) string
	// Info returns the status of the current node.
	Info() string
	// ChangeDirectory switches the currently active node
	ChangeDirectory(path string) string
}
