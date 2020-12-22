package main

import (
	"testing"
)

func TestSearchGetContentLongContent(t *testing.T) {
	var maximum int = 3
	var content string = "abcdef"

	search := NewSearch()
	search.SetActive(true)
	search.AppendStringToContent(content)

	var expected string = "def"
	var result string = search.GetContent(maximum)
	if result != expected {
		t.Errorf("got %s, expected %s", result, expected)
	}
}

func TestSearchGetContentShortContent(t *testing.T) {
	var maximum int = 3
	var content string = "abc"

	search := NewSearch()
	search.SetActive(true)
	search.AppendStringToContent(content)

	var expected string = "abc"
	var result string = search.GetContent(maximum)
	if result != expected {
		t.Errorf("got %s, expected %s", result, expected)
	}
}

func TestSearchActive(t *testing.T) {
	var content string = "foo"

	search := NewSearch()
	search.SetActive(false)
	search.AppendStringToContent(content)

	if search.content != "" {
		t.Errorf("Search added content %s while inactive", search.content)
	}
}

func TestSearchPop(t *testing.T) {
	var content string = "bar"

	search := NewSearch()
	search.SetActive(true)
	search.AppendStringToContent(content)

	var expected string = "ba"
	search.Pop()
	if search.content != expected {
		t.Errorf("got %s, expected %s", search.content, expected)
	}
}

func TestSearchPopEmpty(t *testing.T) {
	search := NewSearch()

	search.Pop()
	if search.content != "" {
		t.Error("pop unsuccessful")
	}
}
