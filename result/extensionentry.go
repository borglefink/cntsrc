package result

// ExtensionEntry is the file entry for result
type ExtensionEntry struct {
	ExtensionName string
	IsBinary      bool
	NumberOfFiles int
	NumberOfLines int
	Filesize      int64
}
