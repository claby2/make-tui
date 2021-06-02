package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Render sets up and renders widgets to build the user interface
func Render(content *ParsedContent, theme string) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	content.Content = replaceTabs(content.Content)
	termWidth, termHeight := ui.TerminalDimensions()

	target := NewTarget(0, len(content.Rules), content.Rules)
	target.Rows = getTargets(content.Rules)

	highlighter := NewHighlighter(theme)

	dependencyWidget := widgets.NewParagraph()
	dependencyWidget.Title = "Dependencies"
	dependencyWidget.Text = getDependency(content.Rules, target.Index)

	contentWidget := widgets.NewParagraph()
	contentWidget.Title = content.FilePath
	contentWidget.Text = getContent(content.Content, highlighter, content.Rules, termHeight, target.Index)

	grid := ui.NewGrid()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewCol(0.2,
			ui.NewRow(0.1, dependencyWidget),
			ui.NewRow(0.9, target),
		),
		ui.NewCol(0.8, contentWidget),
	)

	ui.Render(grid)

	uiEvents := ui.PollEvents()
	previousKey := ""
	quit := false
	run := false
	for !quit && !run {
		e := <-uiEvents
		if target.Search.active {
			// Events if in search mode
			if isLetter(e.ID) {
				target.Search.AppendStringToContent(e.ID)
			} else {
				switch e.ID {
				case "<Backspace>":
					target.Search.Pop()
				case "<Enter>":
					// Search for target
					var index int = target.FindTarget(target.Search.content)
					if index != -1 {
						target.ScrollAmount(index - target.Index)
					}
					fallthrough
				case "<Escape>":
					target.Search.SetActive(false)
				}

			}
		} else {
			// Events if not in search mode
			switch e.ID {
			case "q", "<C-c>":
				quit = true
			case "j", "<Down>":
				target.ScrollDown()
			case "k", "<Up>":
				target.ScrollUp()
			case "g":
				if previousKey == "g" {
					target.ScrollTop()
				}
			case "G", "<End>":
				target.ScrollBottom()
			case "<C-d>":
				target.ScrollHalfPageDown()
			case "<C-u>":
				target.ScrollHalfPageUp()
			case "<C-f>":
				target.ScrollPageDown()
			case "<C-b>":
				target.ScrollPageUp()
			case "/":
				target.Search.SetActive(true)
			case "<Enter>":
				run = true
			}
		}
		if previousKey == e.ID {
			previousKey = ""
		} else {
			previousKey = e.ID
		}
		// Global events
		switch e.ID {
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
			termWidth, termHeight = ui.TerminalDimensions()
			ui.Clear()
		}
		target.Index = target.SelectedRow
		dependencyWidget.Text = getDependency(content.Rules, target.Index)
		contentWidget.Text = getContent(content.Content, highlighter, content.Rules, termHeight, target.Index)
		ui.Render(grid)
	}
	ui.Close()

	if run {
		runTarget(target.GetName(), content.FilePath)
	}
}

func getTargets(rules []Rule) []string {
	// Target names are the first element in each slice in rules
	if len(rules) == 0 {
		// No rules were found
		return []string{""}
	}
	var targets []string
	for _, rule := range rules {
		targets = append(targets, rule.target)
	}
	return targets
}

func getDependency(rules []Rule, index int) string {
	if index < len(rules) {
		return rules[index].dependencies
	}
	return ""
}

func getContent(content []string, highlighter *Highlighter, rules []Rule, termHeight, index int) string {
	contentCopy := append([]string(nil), content...)
	highlightedContent := highlighter.GetHighlightedContent(contentCopy)
	firstLine := 0
	if index < len(rules) {
		lineNumber := rules[index].lineNumber

		if len(content) > termHeight-1 {
			firstLine = lineNumber
		}

		// Highlight rule (including commands)
		highlightedContent[lineNumber] = "[" + content[lineNumber] + "](fg:black,bg:white,mod:bold)"
	}
	return strings.ReplaceAll(strings.Join(highlightedContent[firstLine:], "\n"), "\t", strings.Repeat(" ", 4))
}

func replaceTabs(content []string) []string {
	contentCopy := append([]string(nil), content...)
	for i, line := range contentCopy {
		contentCopy[i] = strings.ReplaceAll(line, "\t", strings.Repeat(" ", 4))
	}
	return contentCopy
}

func isLetter(s string) bool {
	return s[0] != '<' && s[len(s)-1] != '>'
}

func runTarget(target, filePath string) {
	if target != "" {
		// Compose command
		cmd := exec.Command("make", "-f"+filePath, target)
		// CombinedOutput will return both standard output and standard error
		stdoutStderr, _ := cmd.CombinedOutput()
		fmt.Print(string(stdoutStderr))
	}
}
