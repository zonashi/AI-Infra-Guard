// Package utils favicon相关工具
package utils

import (
	"bytes"
	"github.com/Tencent/AI-Infra-Guard/pkg/httpx"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetFaviconBytes retrieves the favicon bytes from a website
// 这个函数用于获取网站的favicon图标数据
func GetFaviconBytes(hp *httpx.HTTPX, domain string, resp []byte) ([]byte, error) {
	// Load the HTML document from response bytes
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp))
	if err != nil {
		return nil, err
	}

	// Construct default favicon.ico URL
	faviconUrl, err := url.JoinPath(domain, "/favicon.ico")
	if err != nil {
		return nil, err
	}

	// Initialize slice to store potential favicon URLs
	var urlList []string = []string{}

	// Find all <link> tags and extract favicon URLs
	// Look for rel attributes containing "icon"
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		rel, ok := s.Attr("rel")
		if ok && strings.Contains(strings.ToLower(rel), "icon") {
			href, ok := s.Attr("href")
			if ok {
				// Join domain with href to create absolute URL
				href, err = url.JoinPath(domain, href)
				if err == nil {
					urlList = append(urlList, href)
				}
			}
		}
	})

	// Add default favicon.ico URL as fallback
	urlList = append(urlList, faviconUrl)

	// Try each URL until we successfully get a favicon
	for _, u := range urlList {
		httpResp, err := hp.Get(u, nil)
		if err != nil {
			continue
		}
		return httpResp.Data, nil
	}

	return nil, nil
}
