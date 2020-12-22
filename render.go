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
func Render(content *ParsedContent) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	content.content = replaceTabs(content.content)
	termWidth, termHeight := ui.TerminalDimensions()

	target := NewTarget(0, len(content.rules), content.rules)
	target.Rows = getTargets(content.rules)

	dependencyWidget := widgets.NewParagraph()
	dependencyWidget.Title = "Dependencies"
	dependencyWidget.Text = getDependency(content.rules, target.index)

	contentWidget := widgets.NewParagraph()
	contentWidget.Title = content.filePath
	contentWidget.Text = getHighlightedContent(content.content, content.rules, termHeight, target.index)

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
		if target.search.active {
			// Events if in search mode
			if isLetter(e.ID) {
				target.search.AppendStringToContent(e.ID)
			} else {
				switch e.ID {
				case "<Backspace>":
					target.search.Pop()
				case "<Enter>":
					// Search for target
					var index int = target.FindTarget(target.search.content)
					if index != -1 {
						target.ScrollAmount(index - target.index)
						target.SetIndex(index)
					}
					fallthrough
				case "<Escape>":
					target.search.SetActive(false)
				}

			}
		} else {
			// Events if not in search mode
			switch e.ID {
			case "q", "<C-c>":
				quit = true
			case "j", "<Down>":
				target.ScrollDown()
				target.Down(1)
			case "k", "<Up>":
				target.ScrollUp()
				target.Up(1)
			case "g":
				if previousKey == "g" {
					target.ScrollTop()
					target.Top()
				}
			case "G", "<End>":
				target.ScrollBottom()
				target.Bottom()
			case "<C-d>":
				target.ScrollHalfPageDown()
				target.HalfPageDown(target.Inner.Dy())
			case "<C-u>":
				target.ScrollHalfPageUp()
				target.HalfPageUp(target.Inner.Dy())
			case "/":
				target.search.SetActive(true)
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

		dependencyWidget.Text = getDependency(content.rules, target.index)
		contentWidget.Text = getHighlightedContent(content.content, content.rules, termHeight, target.index)
		ui.Render(grid)
	}

	ui.Close()
	if run && target.name != "" {
		fmt.Println("make", target.name)
		output, err := exec.Command("make", "-f"+content.filePath, target.name).CombinedOutput()
		fmt.Println(string(output))
		if err != nil {
			log.Fatal(err)
		}
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

func getHighlightedContent(content []string, rules []Rule, termHeight, index int) string {
	contentCopy := append([]string(nil), content...)
	firstLine := 0
	if index < len(rules) {
		lineNumber := rules[index].lineNumber
		numberOfCommands := len(rules[index].commands)

		if len(contentCopy) > termHeight-1 {
			firstLine = lineNumber
		}

		// Highlight rule (including commands)
		for i := lineNumber; i <= lineNumber+numberOfCommands; i++ {
			contentCopy[i] = "[" + contentCopy[i] + "](fg:black,bg:white,mod:bold)"
		}
	}
	return strings.ReplaceAll(strings.Join(contentCopy[firstLine:], "\n"), "\t", strings.Repeat(" ", 4))
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
