package core

type fileBasedFileDataNodeProvider struct {
	access fileAccess
}

func newFileDataNodeProvider(access fileAccess) FileDataNodeProvider {
	provider := &fileBasedFileDataNodeProvider{
		access: access}

	return provider
}

// Provide tries to resolve and return a DataNode for the given file.
func (provider *fileBasedFileDataNodeProvider) Provide(filePath, fileName string, parent DataNode) DataNode {
	return nil
}
