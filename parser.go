package main

import (
	"regexp"
	"strings"
)

// Rule contains information about a rule
type Rule struct {
	target, dependencies string
	commands             []string
	lineNumber           int
}

// NewRule constructs a rule given the target, dependencies, commands, and line number in the Makefile
func NewRule(target, dependencies string, commands []string, lineNumber int) *Rule {
	return &Rule{target: target, dependencies: dependencies, commands: commands, lineNumber: lineNumber}
}

// ParsedContent contains the content of a Makefile with its parsed rules
type ParsedContent struct {
	filePath              string
	includeSpecialTargets bool
	content               []string
	rules                 []Rule
}

// NewParsedContent constructs ParsedContent and stages the Makefile for parsing
func NewParsedContent(filePath string, content []string) *ParsedContent {
	return &ParsedContent{filePath: filePath, includeSpecialTargets: false, content: content, rules: []Rule{}}
}

// SetIncludeSpecialTargets sets the option to include special targets to the given boolean value
func (parsedContent *ParsedContent) SetIncludeSpecialTargets(value bool) {
	parsedContent.includeSpecialTargets = value
}

// Parse parses the content of a Makefile and extracts rules from it
func (parsedContent *ParsedContent) Parse() {
	inMultilineComment := false
	inTarget := false
	multilineCommentRegexp := regexp.MustCompile(`^.*#.*\\$`)
	ruleRegexp := regexp.MustCompile(`^([^:\s]+)\s*:\s*([^=].*)?$`)
	for lineNumber, line := range parsedContent.content {
		// Handle multiline comments
		if inMultilineComment {
			inTarget = false
			inMultilineComment = line[len(line)-1:] == "\\"
			// If currently in multiline comment, the entire line is commented
			continue
		} else {
			// Check if current line is start of multiline comment
			multilineMatch := multilineCommentRegexp.MatchString(line)
			inMultilineComment = multilineMatch
		}

		line := stripComments(line)

		// Handle rule commands
		if inTarget && (len(line) == 0 || line[0] == '\t') {
			// Current line is a command
			ruleIndex := len(parsedContent.rules) - 1
			parsedContent.rules[ruleIndex].commands = append(parsedContent.rules[ruleIndex].commands, strings.TrimSpace(line))
			continue
		} else if inTarget {
			inTarget = false
		}

		ruleSubmatch := ruleRegexp.FindStringSubmatch(line)
		if ruleSubmatch != nil && !(!parsedContent.includeSpecialTargets && isSpecialTarget(ruleSubmatch[1])) {
			// Match has been found
			newRule := NewRule(ruleSubmatch[1], ruleSubmatch[2], []string{}, lineNumber)
			parsedContent.rules = append(parsedContent.rules, *newRule)
			inTarget = true
		}
	}
}

func stripComments(line string) string {
	inQuotes := false
	for index, c := range line {
		if c == '"' {
			inQuotes = !inQuotes
		} else if !inQuotes && c == '#' {
			// Slice string from comment delimiter
			line = strings.TrimSpace(line[:index])
			break
		}
	}
	return line
}

func isSpecialTarget(target string) bool {
	// Returns true if the given target is a special built-in target name
	specialTargetNames := []string{
		".PHONY",
		".SUFFIXES",
		".DEFAULT",
		".PRECIOUS",
		".INTERMEDIATE",
		".SECONDARY",
		".SECONDEXPANSION",
		".DELETE_ON_ERROR",
		".IGNORE",
		".LOW_RESOLUTION_TIME",
		".SILENT",
		".EXPORT_ALL_VARIABLES",
		".NOTPARALLEL",
		".ONESHELL",
		".POSIX",
	}
	for _, specialTarget := range specialTargetNames {
		if target == specialTarget {
			return true
		}
	}
	return false
}
