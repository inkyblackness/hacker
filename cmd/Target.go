package cmd

// A Target is an evaluation target, capabile of processing commands
type Target interface {
	// Load requests to load data files from two paths.
	Load(path1, path2 string) string
}
