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
	numberOfRules := 3
	target := NewTarget(0, numberOfRules, getRules(numberOfRules))

	expected := 2
	result := target.FindTarget("target2")
	if result != expected {
		t.Errorf("got index %d, expected index %d", result, expected)
	}
}

func TestTargetFindFakeTarget(t *testing.T) {
	numberOfRules := 3
	target := NewTarget(0, numberOfRules, getRules(numberOfRules))

	expected := -1
	result := target.FindTarget("fake")
	if result != expected {
		t.Errorf("got index %d, expected index %d", result, expected)
	}
}
