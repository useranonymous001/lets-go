package storage

import "fmt"

type MemoryStorage struct {
	data map[string]string
}

// cretes a space for memory storage
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string]string),
	}
}

func (ms *MemoryStorage) Save(key, value string) error {
	ms.data[key] = value
	return nil
}

func (ms *MemoryStorage) Load(key string) (string, error) {
	if val, ok := ms.data[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("key not found")
}
