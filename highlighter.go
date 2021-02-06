package main

import (
	"log"
	"math"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
)

var c = chroma.MustParseColour

// ColorMap associates termui color names with a respective chroma color
var ColorMap = map[string]chroma.Colour{
	"red":     c("#ff0000"),
	"blue":    c("#0000ff"),
	"black":   c("#000000"),
	"cyan":    c("#00ffff"),
	"yellow":  c("#ffff00"),
	"white":   c("#ffffff"),
	"green":   c("#00ff00"),
	"magenta": c("#ff00ff"),
}

// Highlighter helps facilitate makefile syntax highlighting
type Highlighter struct {
	style *chroma.Style
}

// NewHighlighter constructs a Highlighter and sets the style based on the given styleName
func NewHighlighter(styleName string) *Highlighter {
	style := styles.Get(styleName)
	if style == nil {
		style = styles.Fallback
	}
	return &Highlighter{style}
}

// GetHighlightedContent iterates through the given content slice and returns it with inline style annotations
func (highlighter *Highlighter) GetHighlightedContent(content []string) []string {
	lexer := lexers.Get("Base Makefile")

	var highlightedContent []string
	currentLine := ""
	for _, line := range content {
		iterator, err := lexer.Tokenise(nil, line)
		if err != nil {
			log.Fatal(err)
		}
		for token := iterator(); token != chroma.EOF; token = iterator() {
			if token.Value == "\n" {
				highlightedContent = append(highlightedContent, currentLine)
				currentLine = ""
				continue
			}
			entry := highlighter.style.Get(token.Type)
			fg := "clear"
			if entry.Colour.IsSet() {
				fg = approximateColor(entry.Colour)
			}
			style := "fg:" + fg
			if entry.Bold == chroma.Yes {
				style += ",mod:bold"
			} else if entry.Underline == chroma.Yes {
				style += ",mod:underline"
			}
			currentLine += "[" + token.Value + "](" + style + ")"
		}
		if currentLine != "" {
			highlightedContent = append(highlightedContent, currentLine)
			currentLine = ""
		}
	}
	return highlightedContent
}

func approximateColor(color chroma.Colour) string {
	lowestDistance := math.MaxFloat64
	var bestColor string
	for colorName, c := range ColorMap {
		distance := color.Distance(c)
		if distance < lowestDistance {
			lowestDistance = distance
			bestColor = colorName
		}
	}
	return bestColor
}
