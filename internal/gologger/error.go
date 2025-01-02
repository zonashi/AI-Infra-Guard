// Package gologger error包装器
package gologger

import (
	"fmt"
)

// WarpError combines an existing error with an additional message
// 将现有的错误和新的错误信息组合成一个新的错误
// Parameters:
//   - err: original error object (原始错误对象)
//   - message: additional error message to append (要附加的额外错误信息)
//
// Returns:
//   - error: a new error combining both messages (返回组合后的新错误)
func WarpError(err error, message string) error {
	return fmt.Errorf("%s %s", err.Error(), message)
}
