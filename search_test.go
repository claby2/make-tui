package main

import (
	"testing"
)

func TestSearchGetContent(t *testing.T) {
	var maximum int = 3
	var content string = "abcdef"

	searchManager := NewSearchManager()
	searchManager.SetActive(true)
	searchManager.AppendStringToContent(content)

	var expected string = "def"
	var result string = searchManager.GetContent(maximum)
	if result != expected {
		t.Errorf("got %s, expected %s", result, expected)
	}
}

func TestSearchActive(t *testing.T) {
	var content string = "foo"

	searchManager := NewSearchManager()
	searchManager.SetActive(false)
	searchManager.AppendStringToContent(content)

	if searchManager.content != "" {
		t.Errorf("SearchManager added content %s while inactive", searchManager.content)
	}
}

func TestSearchPop(t *testing.T) {
	var content string = "bar"

	searchManager := NewSearchManager()
	searchManager.SetActive(true)
	searchManager.AppendStringToContent(content)

	var expected string = "ba"
	searchManager.Pop()
	if searchManager.content != expected {
		t.Errorf("got %s, expected %s", searchManager.content, expected)
	}
}

func TestSearchPopEmpty(t *testing.T) {
	searchManager := NewSearchManager()

	searchManager.Pop()
	if searchManager.content != "" {
		t.Error("pop unsuccessful")
	}
}
