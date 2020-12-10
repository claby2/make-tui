package main

import (
	"strconv"
	"testing"
)

func getRules(length int) []Rule {
	var rules []Rule
	for i := 0; i < length; i++ {
		newRule := NewRule("target"+strconv.Itoa(i), "dependency"+strconv.Itoa(i), []string{}, i)
		rules = append(rules, *newRule)

	}
	return rules
}

func TestTargetSetIndex(t *testing.T) {
	var numberOfRules int = 10
	target := NewTarget(0, numberOfRules, getRules(numberOfRules))

	var expected int = 5
	target.SetIndex(expected)
	if target.index != expected {
		t.Errorf("got index %d, expected index %d", target.index, expected)
	}
}

func TestTargetSetIndexOutOfBounds(t *testing.T) {
	var numberOfRules int = 10
	target := NewTarget(0, numberOfRules, getRules(numberOfRules))

	var expected int = numberOfRules - 1
	target.SetIndex(numberOfRules + 1)
	if target.index != expected {
		t.Errorf("got index %d, expected index %d", target.index, expected)
	}
}

func TestTargetSetIndexNegative(t *testing.T) {
	var numberOfRules int = 10
	target := NewTarget(0, numberOfRules, getRules(numberOfRules))

	var expected int = 0
	target.SetIndex(-1)
	if target.index != expected {
		t.Errorf("got index %d, expected index %d", target.index, expected)
	}
}

func TestTargetNoRules(t *testing.T) {
	var numberOfRules int = 0
	target := NewTarget(0, numberOfRules, getRules(numberOfRules))

	var expected int = 0
	target.SetIndex(1)
	if target.index != expected {
		t.Errorf("got index %d, expected index %d", target.index, expected)
	}
}

func TestTargetDown(t *testing.T) {
	var numberOfRules int = 10
	var startingIndex int = 5
	target := NewTarget(startingIndex, numberOfRules, getRules(numberOfRules))

	var delta int = 3
	var expected int = startingIndex + delta
	target.Down(delta)
	if target.index != expected {
		t.Errorf("got index %d, expected index %d", target.index, expected)
	}
}

func TestTargetUp(t *testing.T) {
	var numberOfRules int = 10
	var startingIndex int = 5
	target := NewTarget(startingIndex, numberOfRules, getRules(numberOfRules))

	var delta int = 3
	var expected int = startingIndex - delta
	target.Up(delta)
	if target.index != expected {
		t.Errorf("got index %d, expected index %d", target.index, expected)
	}
}

func TestTargetBottom(t *testing.T) {
	var numberOfRules int = 10
	var startingIndex int = 5
	target := NewTarget(startingIndex, numberOfRules, getRules(numberOfRules))

	var expected int = 9
	target.Bottom()
	if target.index != expected {
		t.Errorf("got index %d, expected index %d", target.index, expected)
	}
}

func TestTargetTop(t *testing.T) {
	var numberOfRules int = 10
	var startingIndex int = 5
	target := NewTarget(startingIndex, numberOfRules, getRules(numberOfRules))

	var expected int = 0
	target.Top()
	if target.index != expected {
		t.Errorf("got index %d, expected index %d", target.index, expected)
	}
}

func TestTargetHalfPageDown(t *testing.T) {
	var numberOfRules int = 10
	var startingIndex int = 2
	target := NewTarget(startingIndex, numberOfRules, getRules(numberOfRules))

	var expected int = 7
	target.HalfPageDown(numberOfRules)
	if target.index != expected {
		t.Errorf("got index %d, expected index %d", target.index, expected)
	}
}

func TestTargetHalfPageUp(t *testing.T) {
	var numberOfRules int = 10
	var startingIndex int = 7
	target := NewTarget(startingIndex, numberOfRules, getRules(numberOfRules))

	var expected int = 2
	target.HalfPageUp(numberOfRules)
	if target.index != expected {
		t.Errorf("got index %d, expected index %d", target.index, expected)
	}
}

func TestTargetFindRealTarget(t *testing.T) {
	var numberOfRules int = 3
	target := NewTarget(0, numberOfRules, getRules(numberOfRules))

	var expected int = 2
	var result int = target.FindTarget("target2")
	if result != expected {
		t.Errorf("got index %d, expected index %d", result, expected)
	}
}

func TestTargetFindFakeTarget(t *testing.T) {
	var numberOfRules int = 3
	target := NewTarget(0, numberOfRules, getRules(numberOfRules))

	var expected int = -1
	var result int = target.FindTarget("fake")
	if result != expected {
		t.Errorf("got index %d, expected index %d", result, expected)
	}
}
