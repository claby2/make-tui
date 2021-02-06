package main

import (
	"testing"
)

func TestGetHighlightedContent(t *testing.T) {
	content := []string{
		"PROJECT = make-tui",
		"BUILD_DIR ?= build",
		"APP_SOURCES = parser.go \\",
		"			  render.go \\",
		"			  search.go \\",
		"			  target.go \\",
		"			  main.go \\",
		"",
		"build: $(APP_SOURCES)",
		"	go build -o $(BUILD_DIR)/make-tui $(APP_SOURCES)",
		".PHONY: build",
		"",
		"run: $(APP_SOURCES)",
		"	go run $(APP_SOURCES)",
		".PHONY: run",
		"",
		"test:",
		"	go test ./...",
		".PHONY: test",
		"",
		"clean:",
		"	rm -rf $(BUILD_DIR)",
		".PHONY: clean",
	}
	expected := len(content)
	highlighter := NewHighlighter("emacs")
	highlightedContent := highlighter.GetHighlightedContent(content)

	result := len(highlightedContent)

	if result != expected {
		t.Errorf("expected highlighted content %d, received %d", expected, result)
	}

}
