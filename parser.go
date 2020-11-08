package main

import (
	"regexp"
	"strings"
)

type ParsedContent struct {
	filePath string
	content  []string
	rules    [][]string
}

func NewParsedContent(newFilePath string, fileContent []string) *ParsedContent {
	return &ParsedContent{filePath: newFilePath, content: fileContent, rules: [][]string{}}
}

func (parsedContent *ParsedContent) Parse() {
	inMultilineComment := false
	multilineCommentRegexp := regexp.MustCompile(`^.*#.*\\$`)
	ruleRegexp := regexp.MustCompile(`^([^:\s]+)\s*:\s*([^=].*)?$`)
	for _, line := range parsedContent.content {
		if inMultilineComment {
			inMultilineComment = line[len(line)-1:] == "\\"
			// If currently in multiline comment, the entire line is commented
			continue
		} else {
			// Check if current line is start of multiline comment
			multilineMatch := multilineCommentRegexp.MatchString(line)
			inMultilineComment = multilineMatch
		}
		line := stripComments(line)
		ruleSubmatch := ruleRegexp.FindStringSubmatch(line)
		if ruleSubmatch != nil {
			// Match has been found
			parsedContent.rules = append(parsedContent.rules, ruleSubmatch[1:])
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
