package main

import (
	"image"
	"strings"

	ui "github.com/gizak/termui/v3"
)

// Search manages target search functionality and rendering
type Search struct {
	ui.Block

	active  bool
	content string
}

// NewSearch constructs a Search and sets active to false and content to an empty strin`
func NewSearch() *Search {
	return &Search{active: false, content: ""}
}

// SetActive sets the Search to be active or inactive
func (search *Search) SetActive(value bool) {
	if !value {
		search.content = ""
	}
	search.active = value
}

// AppendStringToContent appends a given string to content if the Search is active
func (search *Search) AppendStringToContent(str string) {
	if search.active {
		search.content += str
	}
}

// Pop removes the last character of Search content if it is not empty
func (search *Search) Pop() {
	var contentLength int = len(search.content)
	if contentLength > 0 {
		search.content = search.content[:contentLength-1]
	}
}

// GetContent returns the content relative to the given maximum number of characters the search bar can render
func (search *Search) GetContent(maximum int) string {
	var contentLength int = len(search.content)
	if contentLength > maximum {
		return search.content[contentLength-maximum:]
	}
	return search.content
}

// Draw creates and renders the filter string to the buffer based on content
func (search *Search) Draw(buf *ui.Buffer) {
	if !search.active {
		return
	}

	style := ui.NewStyle(ui.ColorWhite, ui.ColorClear)
	cursorStyle := ui.NewStyle(ui.ColorWhite, ui.ColorWhite)
	voidStyle := ui.NewStyle(ui.ColorBlack, ui.ColorBlack)

	label := "Search: ["

	maximumContentLength := search.Max.X - search.Min.X - len(label) - 2

	if maximumContentLength <= 0 {
		return
	}
	// Subtract 1 to take into account cursor
	content := search.GetContent(maximumContentLength)

	// Fill with initial content including start of search block
	label += content
	p := image.Pt(search.Min.X, search.Min.Y)
	buf.SetString(label, style, p)
	p.X += len(label)

	// Render cursor
	buf.SetString(" ", cursorStyle, p)
	p.X++

	// Potentially fill the rest of the search block
	remaining := maximumContentLength - len(content)
	buf.SetString(strings.Repeat(" ", remaining), voidStyle, p)
	p.X += remaining

	// Render end of search block
	buf.SetString("]", style, p)
}
