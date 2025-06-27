package search

type SearchProvider interface {
	Search(string) map[string]string
}
