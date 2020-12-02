package main

import (
	"fmt"
	"log"
	"math"
	"os/exec"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

type Target struct {
	index, numberOfRules int
	name                 string
	targets              []string
}

func NewTarget(index, numberOfRules int, rules []Rule) *Target {
	var targets []string
	for _, rule := range rules {
		targets = append(targets, rule.target)
	}
	var name string
	if len(targets) > 0 {
		name = targets[0]
	}
	return &Target{index: index, numberOfRules: numberOfRules, name: name, targets: targets}
}

func (target *Target) Down(delta int) {
	if target.numberOfRules > 0 {
		target.index = int(math.Min(float64(target.numberOfRules-1), float64(target.index+delta)))
		target.name = target.targets[target.index]
	}
}

func (target *Target) Up(delta int) {
	if target.numberOfRules > 0 {
		target.index = int(math.Max(float64(0), float64(target.index-delta)))
		target.name = target.targets[target.index]
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

func Render(content *ParsedContent) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}

	content.content = replaceTabs(content.content)
	target := NewTarget(0, len(content.rules), content.rules)
	termWidth, termHeight := ui.TerminalDimensions()

	targetsWidget := widgets.NewList()
	targetsWidget.Title = "Targets"
	targetsWidget.Rows = getTargets(content.rules)
	targetsWidget.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)

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
			ui.NewRow(0.8, targetsWidget),
			ui.NewRow(0.2, dependencyWidget),
		),
		ui.NewCol(0.8, contentWidget),
	)
	ui.Render(grid)

	uiEvents := ui.PollEvents()
	quit := false
	run := false
	for !quit && !run {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			quit = true
		case "j", "<Down>":
			targetsWidget.ScrollDown()
			target.Down(1)
		case "k", "<Up>":
			targetsWidget.ScrollUp()
			target.Up(1)
		case "<Enter>":
			run = true
		case "<Resize>":
			payload := e.Payload.(ui.Resize)
			grid.SetRect(0, 0, payload.Width, payload.Height)
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
