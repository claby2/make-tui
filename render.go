package main

import (
	"log"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func GetTargets(rules []Rule) []string {
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

func GetDependency(rules []Rule, index int) string {
	if index < len(rules) {
		return rules[index].dependencies
	}
	return ""
}

func Render(content *ParsedContent) {
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	targetsWidget := widgets.NewList()
	targetsWidget.Title = "Targets"
	targetsWidget.Rows = GetTargets(content.rules)
	targetsWidget.SelectedRowStyle = ui.NewStyle(ui.ColorYellow)

	dependenciesWidget := widgets.NewParagraph()
	dependenciesWidget.Title = "Dependencies"

	contentWidget := widgets.NewParagraph()
	contentWidget.Title = content.filePath
	contentWidget.Text = strings.ReplaceAll(strings.Join(content.content, "\n"), "\t", strings.Repeat(" ", 4))

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewCol(0.2,
			ui.NewRow(0.8, targetsWidget),
			ui.NewRow(0.2, dependenciesWidget),
		),
		ui.NewCol(0.8, contentWidget),
	)
	ui.Render(grid)

	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			return
		case "j", "<Down>":
			targetsWidget.ScrollDown()
		case "k", "<Up>":
			targetsWidget.ScrollUp()
		case "<C-d>":
			targetsWidget.ScrollHalfPageDown()
		case "<C-u>":
			targetsWidget.ScrollHalfPageUp()
		case "<C-f>":
			targetsWidget.ScrollPageDown()
		case "<C-b>":
			targetsWidget.ScrollPageUp()
		}
		dependenciesWidget.Text = GetDependency(content.rules, targetsWidget.SelectedRow)
		ui.Render(grid)
	}
}
