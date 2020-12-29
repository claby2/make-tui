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
