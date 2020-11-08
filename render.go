package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func GetTargets(rules [][]string) []string {
	// Target names are the first element in each slice in rules
	var targets []string
	for _, rule := range rules {
		targets = append(targets, rule[0])
	}
	return targets
}

func GetDependency(rules [][]string, index int) string {
	return rules[index][1]
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

	grid := ui.NewGrid()
	termWidth, termHeight := ui.TerminalDimensions()
	grid.SetRect(0, 0, termWidth, termHeight)
	grid.Set(
		ui.NewCol(0.2,
			ui.NewRow(0.8, targetsWidget),
			ui.NewRow(0.2, dependenciesWidget),
		),
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
