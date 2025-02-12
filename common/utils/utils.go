// Package utils 工具集合
package utils

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/spaolacci/murmur3"
)

// Duration2String 将时间段转换为可读的字符串格式
// 如果时间超过60秒则返回分钟，否则返回秒
func Duration2String(t time.Duration) string {
	sceond := t.Seconds()
	if sceond >= 60 {
		return fmt.Sprintf("%.2f min", t.Minutes())
	} else {
		return fmt.Sprintf("%.2f s", sceond)
	}
}

// InsertInto 在字符串中每隔指定间隔插入分隔符
// s: 源字符串
// interval: 插入间隔
// sep: 分隔符
func InsertInto(s string, interval int, sep rune) string {
	var buffer bytes.Buffer
	before := interval - 1
	last := len(s) - 1
	for i, char := range s {
		buffer.WriteRune(char)
		if i%interval == before && i != last {
			buffer.WriteRune(sep)
		}
	}
	buffer.WriteRune(sep)
	return buffer.String()
}

// FaviconHash 计算网站图标的哈希值
// 将数据进行base64编码后使用murmur3哈希算法计算
func FaviconHash(data []byte) int32 {
	stdBase64 := base64.StdEncoding.EncodeToString(data)
	stdBase64 = InsertInto(stdBase64, 76, '\n')
	hasher := murmur3.New32WithSeed(0)
	hasher.Write([]byte(stdBase64))
	return int32(hasher.Sum32())
}

// ScanDir 递归扫描目录，返回所有文件的完整路径
// path: 要扫描的目录路径
// 返回文件路径列表和可能的错误
func ScanDir(path string) ([]string, error) {
	files := make([]string, 0)
	dir, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, fi := range dir {
		if fi.IsDir() {
			newDir, err := ScanDir(filepath.Join(path, fi.Name()))
			if err != nil {
				return files, err
			}
			files = append(files, newDir...)
		} else {
			files = append(files, filepath.Join(path, fi.Name()))
		}
	}
	return files, nil
}

// IsCIDR 检查给定的字符串是否为有效的CIDR格式
func IsCIDR(cidr string) bool {
	_, _, err := net.ParseCIDR(cidr)
	return err == nil
}

// IsFileExists 检查文件是否存在
// path: 文件路径
// 返回布尔值表示文件是否存在
func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// IsDir 判断给定路径是否为目录
// path: 待检查的路径
// 返回布尔值表示是否为目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// TrimProtocol 移除URL中的HTTP/HTTPS协议前缀
// targetURL: 目标URL
// 返回去除协议前缀后的URL
func TrimProtocol(targetURL string) string {
	URL := strings.TrimSpace(targetURL)
	if strings.HasPrefix(strings.ToLower(URL), "http://") || strings.HasPrefix(strings.ToLower(URL), "https://") {
		URL = URL[strings.Index(URL, "//")+2:]
	}
	URL = strings.TrimRight(URL, "/")
	return URL
}

// CompareVersions 比较两个版本号字符串
// version1, version2: 待比较的版本号
// 返回值: 1 表示 version1 大于 version2
//
//	-1 表示 version1 小于 version2
//	 0 表示两个版本号相等
func CompareVersions(version1, version2 string) int {
	v1Parts := strings.Split(version1, ".")
	v2Parts := strings.Split(version2, ".")

	// Determine the max length to iterate over
	maxLen := len(v1Parts)
	if len(v2Parts) > maxLen {
		maxLen = len(v2Parts)
	}

	for i := 0; i < maxLen; i++ {
		var num1, num2 int

		if i < len(v1Parts) {
			num1, _ = strconv.Atoi(v1Parts[i])
		}

		if i < len(v2Parts) {
			num2, _ = strconv.Atoi(v2Parts[i])
		}

		if num1 > num2 {
			return 1
		} else if num1 < num2 {
			return -1
		}
	}

	return 0
}

// GetMiddleText 获取两个字符串之间的文本内容
// left: 左边界字符串
// right: 右边界字符串
// html: 源文本
// 返回左右边界之间的文本，如果未找到则返回空字符串
func GetMiddleText(left, right, html string) string {
	start := strings.Index(html, left)
	if start == -1 {
		return "" // 如果找不到 left，返回空字符串
	}
	start += len(left)

	end := strings.Index(html[start:], right)
	if end == -1 {
		return "" // 如果找不到 right，返回空字符串
	}
	end += start

	return html[start:end]
}

// PortInfo 存储端口和地址信息
type PortInfo struct {
	Port    int
	Address string
}

// GetLocalOpenPorts 获取本地开放的端口及其地址信息
func GetLocalOpenPorts() ([]PortInfo, error) {
	var portInfos []PortInfo
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("netstat", "-an")
		output, err := cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("执行netstat命令失败: %v", err)
		}

		scanner := bufio.NewScanner(strings.NewReader(string(output)))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "LISTENING") {
				parts := strings.Fields(line)
				if len(parts) >= 2 {
					addrPort := strings.Split(parts[1], ":")
					if len(addrPort) == 2 {
						port, err := strconv.Atoi(addrPort[1])
						if err == nil {
							addr := addrPort[0]
							portInfos = append(portInfos, PortInfo{
								Port:    port,
								Address: addr,
							})
						}
					}
				}
			}
		}

	case "darwin", "linux":
		cmd := exec.Command("lsof", "-i", "-P", "-n")
		output, err := cmd.Output()
		if err != nil {
			return nil, fmt.Errorf("执行lsof命令失败: %v", err)
		}

		scanner := bufio.NewScanner(strings.NewReader(string(output)))
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "LISTEN") {
				parts := strings.Fields(line)
				for _, part := range parts {
					if strings.Contains(part, ":") {
						addrPort := strings.Split(part, ":")
						if len(addrPort) == 2 {
							port, err := strconv.Atoi(addrPort[1])
							if err == nil {
								addr := addrPort[0]
								if addr == "*" || addr == "0.0.0.0" {
									addr = "0.0.0.0"
								} else if addr == "127.0.0.1" || addr == "localhost" {
									addr = "127.0.0.1"
								}
								portInfos = append(portInfos, PortInfo{
									Port:    port,
									Address: addr,
								})
							}
						}
					}
				}
			}
		}

	default:
		return nil, fmt.Errorf("不支持的操作系统: %s", runtime.GOOS)
	}

	// 去重
	seen := make(map[string]bool)
	var result []PortInfo
	for _, info := range portInfos {
		key := fmt.Sprintf("%s:%d", info.Address, info.Port)
		if !seen[key] {
			seen[key] = true
			result = append(result, info)
		}
	}

	return result, nil
}
