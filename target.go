package main

import (
	"math"
)

// Target contains information about the rules of a Makefile and keeps track of the currently selected rule
type Target struct {
	index, numberOfRules int
	name                 string
	targets              []string
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
	return &Target{index: index, numberOfRules: numberOfRules, name: name, targets: targets}
}

// Down increases the index of the target while taking into account the total number of rules, effectively scrolling down the list of targets
func (target *Target) Down(delta int) {
	target.SetIndex(target.index + delta)
}

// Up decreases the index of the target while taking into account the total number of rules, effectively scrolling up the list of targets
func (target *Target) Up(delta int) {
	target.SetIndex(target.index - delta)
}

// Bottom sets the current index to the last target
func (target *Target) Bottom() {
	target.SetIndex(target.numberOfRules - 1)
}

// Top sets the current index to the first target
func (target *Target) Top() {
	target.SetIndex(0)
}

// HalfPageDown moves down an equivalent of half the number of targets
func (target *Target) HalfPageDown(listHeight int) {
	target.SetIndex(target.index + int(math.Floor(float64(listHeight)/2)))
}

// HalfPageUp moves up an equivalent of half the number of targets
func (target *Target) HalfPageUp(listHeight int) {
	target.SetIndex(target.index - int(math.Floor(float64(listHeight)/2)))
}

// SetIndex sets the current index of Target while taking into account bounds
func (target *Target) SetIndex(index int) {
	if target.numberOfRules > 0 {
		target.index = int(math.Max(float64(0), math.Min(float64(target.numberOfRules-1), float64(index))))
		target.name = target.targets[target.index]
	}
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
