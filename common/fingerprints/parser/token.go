// Package parser 实现词法分析
package parser

import (
	"errors"
	"strings"
)

// Token represents a lexical unit in the expression parsing
// 表示表达式解析中的词法单元
type Token struct {
	name    string // token type name
	content string // actual content of the token
}

// Constants defining different types of tokens
const (
	// Content type tokens
	tokenBody   = "body"   // matches body content
	tokenHeader = "header" // matches HTTP headers
	tokenIcon   = "icon"   // matches icon content
	tokenText   = "text"   // matches text content

	// Comparison operators
	tokenContains   = "="  // contains operator
	tokenFullEqual  = "==" // exact match operator
	tokenNotEqual   = "!=" // not equal operator
	tokenRegexEqual = "~=" // regex match operator

	// Logical operators
	tokenAnd = "&&" // logical AND
	tokenOr  = "||" // logical OR

	// Parentheses
	tokenLeftBracket  = "("
	tokenRightBracket = ")"
)

// Version comparison related tokens
const (
	tokenVersion = "version" // version identifier
	tokenGt      = ">"       // greater than
	tokenLt      = "<"       // less than
)

// ParseTokens converts input string to token sequence, supporting text content(quoted),
// comparison ops(=,==,!=,~=), logical ops(&&,||), parentheses and keywords(body,header,icon)
func ParseTokens(s1 string) ([]Token, error) {
	return parseTokensWithOptions(s1, []string{tokenBody, tokenHeader, tokenIcon})
}

// ParseAdvisorTokens parses advisor expressions, similar to ParseTokens but supports version keyword
func ParseAdvisorTokens(s1 string) ([]Token, error) {
	return parseTokensWithOptions(s1, []string{tokenVersion})
}

// parseTokensWithOptions 是提取的公共解析函数
func parseTokensWithOptions(s1 string, validKeywords []string) ([]Token, error) {
	s, tokens := []rune(s1), []Token{}
	for i := 0; i < len(s); {
		switch x := s[i]; x {
		case '"':
			token, newPos, err := parseQuotedText(s, i)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			i = newPos + 1
		case '=', '~', '!', '|', '&':
			token, skip := parseOperator(s[i:])
			if token.name != "" {
				tokens = append(tokens, token)
				i += skip
			}
		case '(', ')':
			tokens = append(tokens, Token{
				name:    map[rune]string{'(': tokenLeftBracket, ')': tokenRightBracket}[x],
				content: string(x),
			})
			i++
		case '>', '<':
			tokens = append(tokens, Token{
				name:    map[rune]string{'>': tokenGt, '<': tokenLt}[x],
				content: string(x),
			})
			i++
		case ' ':
			i++
		default:
			token, newPos, err := parseKeyword(s[i:], validKeywords)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			i += newPos
		}
	}
	return tokens, nil
}

// 辅助函数：解析引号内的文本
func parseQuotedText(s []rune, start int) (Token, int, error) {
	var n []rune
	i := start + 1
	for i < len(s) {
		if s[i] == '\\' {
			n = append(n, s[i+1])
			i += 2
		} else if s[i] == '"' {
			return Token{name: tokenText, content: string(n)}, i, nil
		} else {
			n = append(n, s[i])
			i++
		}
	}
	return Token{}, 0, errors.New("unknown text:" + string(s[start:]))
}

// 辅助函数：解析操作符
func parseOperator(s []rune) (Token, int) {
	ops := map[string]struct {
		name, content string
		skip          int
	}{
		"==": {tokenFullEqual, "==", 2},
		"=":  {tokenContains, "=", 1},
		"~=": {tokenRegexEqual, "~=", 2},
		"!=": {tokenNotEqual, "!=", 2},
		"||": {tokenOr, "||", 2},
		"&&": {tokenAnd, "&&", 2},
	}
	for op, info := range ops {
		if strings.HasPrefix(string(s), op) {
			return Token{name: info.name, content: info.content}, info.skip
		}
	}
	return Token{}, 0
}

// CheckBalance verifies if parentheses in token sequence are balanced
// Returns error if unbalanced, nil otherwise
// 主要功能:检查token序列中的括号是否匹配
// 不匹配时返回error,匹配时返回nil
func CheckBalance(tokens []Token) error {
	stream := newTokenStream(tokens)
	var parens int
	for stream.hasNext() {
		tmpToken, err := stream.next()
		if err != nil {
			return err
		}
		if tmpToken.name == tokenLeftBracket {
			parens++
			continue
		}
		if tmpToken.name == tokenRightBracket {
			parens--
			continue
		}
	}
	if parens != 0 {
		return errors.New("unbalanced parenthesis")
	}
	return nil
}

// 辅助函数：解析关键字
func parseKeyword(s []rune, validKeywords []string) (Token, int, error) {
	textOption := string(s)
	for _, check := range validKeywords {
		if strings.HasPrefix(textOption, check) {
			return Token{
				name:    check,
				content: check,
			}, len(check), nil
		}
	}
	return Token{}, 0, errors.New("unknown text:" + textOption)
}
