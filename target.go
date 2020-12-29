package main

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Target contains information about and renders the rules of a Makefile
type Target struct {
	*widgets.List

	Search        *Search
	Index         int
	Name          string
	targets       []string
	numberOfRules int
}

// NewTarget constructs a Target and decomposes rules into individual target strings
func NewTarget(index, numberOfRules int, rules []Rule) *Target {
	var targets []string
	for _, rule := range rules {
		targets = append(targets, rule.target)
	}
	var name string
	if len(targets) > 0 {
		name = targets[0]
	}
	target := &Target{
		List:          widgets.NewList(),
		Index:         index,
		numberOfRules: numberOfRules,
		Name:          name,
		targets:       targets,
	}

	target.Title = "Targets"
	target.SelectedRowStyle = ui.NewStyle(ui.ColorBlack, ui.ColorWhite)

	//target.search = NewSearch()
	target.Search = &Search{
		active:  false,
		content: "",
	}

	return target
}

// FindTarget returns the index of the target with the given name, returns -1 if it does not exist
func (target *Target) FindTarget(goalTargetName string) int {
	for i, targetName := range target.targets {
		if targetName == goalTargetName {
			return i
		}
	}
	return -1
}

// SetRect sets the rectangle for the target widget for rendering
func (target *Target) SetRect(x1, y1, x2, y2 int) {
	target.List.SetRect(x1, y1, x2, y2)
	// Position search at the bottom
	target.Search.SetRect(x1+2, y2-1, x2-2, y2)
}

// Draw draws the search and target widgets
func (target *Target) Draw(buf *ui.Buffer) {
	target.List.Draw(buf)
	target.Search.Draw(buf)
}
