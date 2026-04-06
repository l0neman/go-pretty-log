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

// GetLine 获取高亮突出显示的一行
func GetLine(text string, width int) string {
	content := "┃ " + text
	return getHighlightLine(content, width)
}

// GetLines 获取高亮突出显示的若干行
func GetLines(texts []string, width int) string {
	content := ""
	for _, text := range texts {
		content += "┃ " + text + "\n"
	}

	content = strings.TrimSuffix(content, "\n")

	return getHighlightLine(content, width)
}
