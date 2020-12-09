package main

// SearchManager manages target search functionality
type SearchManager struct {
	active  bool
	content string
}

// NewSearchManager constructs a SearchManager and sets active to false and content to an empty strin`
func NewSearchManager() *SearchManager {
	return &SearchManager{active: false, content: ""}
}

// SetActive sets the SearchManager to be active or inactive
func (searchManager *SearchManager) SetActive(value bool) {
	if !value {
		searchManager.content = ""
	}
	searchManager.active = value
}

// AppendStringToContent appends a given string to content if the SearchManager is active
func (searchManager *SearchManager) AppendStringToContent(str string) {
	if searchManager.active {
		searchManager.content += str
	}
}

// Pop removes the last character of SearchManager content if it is not empty
func (searchManager *SearchManager) Pop() {
	var contentLength int = len(searchManager.content)
	if contentLength > 0 {
		searchManager.content = searchManager.content[:contentLength-1]
	}
}

// GetContent returns the content relative to the given maximum number of characters the search bar can render
func (searchManager *SearchManager) GetContent(maximum int) string {
	var contentLength int = len(searchManager.content)
	if contentLength > maximum {
		return searchManager.content[contentLength-maximum:]
	}
	return searchManager.content
}
