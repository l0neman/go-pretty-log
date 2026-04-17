package highlignt

import "strings"

func getHighlightLine(content string, width int) string {
	// 1. top line
	line := "┏"
	for i := 0; i < width-1; i++ {
		line += "━"
	}

	// 2. text
	line += "\n"
	line += content
	line += "\n"

	// 3. bottom line
	line += "┗"
	for i := 0; i < width-1; i++ {
		line += "━"
	}

	return line
}

// GetLine returns a highlighted single line
func GetLine(text string, width int) string {
	content := "┃ " + text
	return getHighlightLine(content, width)
}

// GetLines returns highlighted multiple lines
func GetLines(texts []string, width int) string {
	content := ""
	for _, text := range texts {
		content += "┃ " + text + "\n"
	}

	content = strings.TrimSuffix(content, "\n")

	return getHighlightLine(content, width)
}
