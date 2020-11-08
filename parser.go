package main

import (
	"regexp"
	"strings"
)

type Rule struct {
	target, dependencies string
	commands             []string
	lineNumber           int
}

func NewRule(target, dependencies string, commands []string, lineNumber int) *Rule {
	return &Rule{target: target, dependencies: dependencies, commands: commands, lineNumber: lineNumber}
}

type ParsedContent struct {
	filePath string
	content  []string
	rules    []Rule
}

func NewParsedContent(filePath string, content []string) *ParsedContent {
	return &ParsedContent{filePath: filePath, content: content, rules: []Rule{}}
}

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
		if ruleSubmatch != nil {
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
