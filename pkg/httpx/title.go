// Package httpx title
package httpx

import (
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

// ExtractTitle from a response
func ExtractTitle(resp string) (title string) {
	var re = regexp.MustCompile(`(?im)<\s*title.*>(.*?)<\s*/\s*title>`)
	for _, match := range re.FindAllString(resp, -1) {
		title = html.UnescapeString(trimTitleTags(match))
		break
	}
	return
}

func trimTitleTags(title string) string {
	// trim <title>*</title>
	titleBegin := strings.Index(title, ">")
	titleEnd := strings.Index(title, "</")
	return title[titleBegin+1 : titleEnd]
}
