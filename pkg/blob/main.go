package blob

type BlobReader struct {
	Content []byte
	Size    int64
	Name    string
}

type BlobStore interface {
	// GetString returns the string value of the given key.
	Save(key *string, content *[]byte, overwrite bool) error
	Exists(key *string) bool
	Read(key *string, process *Process) (*BlobReader, error)
}
