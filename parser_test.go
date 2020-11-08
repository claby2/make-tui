package main

import (
	"testing"
)

func assertParsedEqual(t *testing.T, result, expected [][]string) {
	if (result == nil) != (expected == nil) || len(result) != len(expected) {
		t.Error("result and expected do not match")
		t.Logf("got %v, expected %v", result, expected)
		return
	}
	for i := range result {
		pass := true
		for j := range result {
			if result[i][j] != expected[i][j] {
				pass = false
				break
			}
		}
		if !pass {
			t.Errorf("got %s, expected %s", result[i], expected[i])
			return
		}
	}
}

func assertParsed(t *testing.T, fileContent []string, expected [][]string) {
	content := NewParsedContent("fileName", fileContent)
	content.Parse()
	assertParsedEqual(t, content.rules, expected)
}

func TestParserNoDependencies(t *testing.T) {
	fileContent := []string{
		"target_1:",
		"\tcommand(s)",
		"target_2:",
	}
	expected := [][]string{
		{"target_1", ""},
		{"target_2", ""},
	}
	assertParsed(t, fileContent, expected)
}

func TestParserComments(t *testing.T) {
	fileContent := []string{
		"#target_1: dependencies_1",
		"\t# target_2: dependencies_2",
		"target_3: dependencies_3 # comment",
	}
	expected := [][]string{
		{"target_3", "dependencies_3"},
	}
	assertParsed(t, fileContent, expected)
}

func TestParserIgnoreCommentInQuotes(t *testing.T) {
	fileContent := []string{
		"target_1: dependencies_1 \"# in quotes\"",
	}
	expected := [][]string{
		{"target_1", "dependencies_1 \"# in quotes\""},
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
	expected := [][]string{
		{"target_1", "dependencies_1"},
		{"target_4", "dependencies_4"},
	}
	assertParsed(t, fileContent, expected)
}
