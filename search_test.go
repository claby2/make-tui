package main

import (
	"testing"
)

func TestSearchGetContentLongContent(t *testing.T) {
	maximum := 3
	content := "abcdef"

	search := NewSearch()
	search.SetActive(true)
	search.AppendStringToContent(content)

	expected := "def"
	result := search.GetContent(maximum)
	if result != expected {
		t.Errorf("got %s, expected %s", result, expected)
	}
}

func TestSearchGetContentShortContent(t *testing.T) {
	maximum := 3
	content := "abc"

	search := NewSearch()
	search.SetActive(true)
	search.AppendStringToContent(content)

	expected := "abc"
	result := search.GetContent(maximum)
	if result != expected {
		t.Errorf("got %s, expected %s", result, expected)
	}
}

func TestSearchActive(t *testing.T) {
	content := "foo"

	search := NewSearch()
	search.SetActive(false)
	search.AppendStringToContent(content)

	if search.content != "" {
		t.Errorf("Search added content %s while inactive", search.content)
	}
}

func TestSearchPop(t *testing.T) {
	content := "bar"

	search := NewSearch()
	search.SetActive(true)
	search.AppendStringToContent(content)

	expected := "ba"
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
