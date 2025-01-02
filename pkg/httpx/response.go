// Package httpx http response
package httpx

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/projectdiscovery/rawhttp/client"
)

// Response contains the response to a server
type Response struct {
	*http.Response
	StatusCode    int
	Headers       map[string][]string
	Data          []byte
	DataStr       string
	ContentLength int
	Title         string
}

// GetHeaderRaw 获得header文本
func (r *Response) GetHeaderRaw() string {
	HeaderStr := ""
	for h, v := range r.Headers {
		HeaderStr += fmt.Sprintf("%s: %s\n", h, strings.Join(v, " "))
	}
	return HeaderStr
}

// GetHeader value
func (r *Response) GetHeader(name string) string {
	v, ok := r.Headers[name]
	if ok {
		return strings.Join(v, " ")
	}
	return ""
}

// GetHeaderPart with offset
func (r *Response) GetHeaderPart(name, sep string) string {
	v, ok := r.Headers[name]
	if ok && len(v) > 0 {
		tokens := strings.Split(strings.Join(v, " "), sep)
		return tokens[0]
	}
	return ""
}

// DumpResponse 导出返回包
func (r *Response) DumpResponse() string {
	firstLine := r.Response.Proto + " " + r.Response.Status
	return firstLine + client.NewLine + r.GetHeaderRaw() + client.NewLine + r.DataStr
}
