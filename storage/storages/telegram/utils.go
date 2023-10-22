package telegram

import (
	"regexp"
)

func escapeMarkdown(text string) string {
	escapeChars := `\_*[]()~` + "`" + ">#+-=|{}.!"
	re := regexp.MustCompile("([" + regexp.QuoteMeta(escapeChars) + "])")
	return re.ReplaceAllString(text, "\\$1")
}
