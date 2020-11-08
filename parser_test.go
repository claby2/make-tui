package main

import (
	"testing"
)

func isRuleEqual(ruleA, ruleB Rule) bool {
	if ruleA.target != ruleB.target ||
		ruleA.dependencies != ruleB.dependencies ||
		(ruleA.commands == nil) != (ruleB.commands == nil) ||
		len(ruleA.commands) != len(ruleB.commands) {
		return false
	}
	for i := range ruleA.commands {
		if ruleA.commands[i] != ruleB.commands[i] {
			return false
		}
	}
	return true
}

func assertRulesEqual(t *testing.T, result, expected []Rule) {
	if (result == nil) != (expected == nil) || len(result) != len(expected) {
		t.Error("result and expected do not match")
		t.Logf("got %v, expected %v", result, expected)
		return
	}
	for i := range result {
		if !isRuleEqual(result[i], expected[i]) {
			t.Errorf("got %v, expected %v", result[i], expected[i])
		}
	}
}

func assertParsed(t *testing.T, fileContent []string, expected []Rule) {
	content := NewParsedContent("fileName", fileContent)
	content.Parse()
	assertRulesEqual(t, content.rules, expected)
}

func TestParserNoDependencies(t *testing.T) {
	fileContent := []string{
		"target_1:",
		"target_2:",
	}
	expected := []Rule{
		{"target_1", "", []string{}},
		{"target_2", "", []string{}},
	}
	assertParsed(t, fileContent, expected)
}

func TestParserWithCommands(t *testing.T) {
	fileContent := []string{
		"target_1: dependencies_1",
		"\tcommand_1",
		"\t\tcommand_2",
		"",
		"\tcommand_3",
		"not_command",
		"target_2:",
		"not_command",
		"not_target",
		"\tnot_command",
	}
	expected := []Rule{
		{"target_1", "dependencies_1", []string{"command_1", "command_2", "", "command_3"}},
		{"target_2", "", []string{}},
	}
	assertParsed(t, fileContent, expected)
}

func TestParserComments(t *testing.T) {
	fileContent := []string{
		"#target_1: dependencies_1",
		"\t# target_2: dependencies_2",
		"target_3: dependencies_3 # comment",
	}
	expected := []Rule{
		{"target_3", "dependencies_3", []string{}},
	}
	assertParsed(t, fileContent, expected)
}

func TestParserIgnoreCommentInQuotes(t *testing.T) {
	fileContent := []string{
		"target_1: dependencies_1 \"# in quotes\"",
	}
	expected := []Rule{
		{"target_1", "dependencies_1 \"# in quotes\"", []string{}},
	}
	assertParsed(t, fileContent, expected)
}

func TestParserMultilineComments(t *testing.T) {
	fileContent := []string{
		"target_1: dependencies_1 # comment \\",
		"target_2: dependencies_2 \\",
		"target_3: dependencies_3 \\ comment",
		"target_4: dependencies_4",
		"# comment \\",
		"target_5: dependencies_5",
	}
	expected := []Rule{
		{"target_1", "dependencies_1", []string{}},
		{"target_4", "dependencies_4", []string{""}},
	}
	assertParsed(t, fileContent, expected)
}
