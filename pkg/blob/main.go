package blob

type BlobStore interface {
	// GetString returns the string value of the given key.
	Save(key *string, content *[]byte, overwrite bool) error
	Exists(key *string) bool
}
