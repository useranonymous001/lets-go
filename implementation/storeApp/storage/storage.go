package storage

// interface is defined here for all kind of storage devices

type Storage interface {
	Save(key, value string) error    // returns error if occured while saving
	Load(key string) (string, error) // returns data and error while reading
}
