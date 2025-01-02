// Package hunyuan 大模型接口实现
package hunyuan

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

// HunyuanAI Golang实现
func HunyuanAI(prompt string, key string) (string, error) {
	// 设置默认值
	if key == "" {
		key = os.Getenv("OPENAI_API_KEY")
	}
	// 检查key是否为空
	if key == "" {
		return "", errors.New("OPENAI_API_KEY is empty")
	}
	baseUrl := os.Getenv("OPENAI_BASE_URL") // 默认值
	if baseUrl == "" {
		baseUrl = "https://hunyuan.tencentcloudapi.com"
	}
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = "hunyuan-turbo"
	}
	client := openai.NewClient(option.WithBaseURL(baseUrl), option.WithAPIKey(key))
	ctx := context.Background()
	stream := client.Chat.Completions.NewStreaming(ctx, openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		}),
		Seed:  openai.Int(24),
		Model: openai.F(model),
	})
	s := strings.Builder{}
	// 循环读取结果
	for stream.Next() {
		evt := stream.Current()
		if len(evt.Choices) > 0 {
			word := evt.Choices[0].Delta.Content
			s.WriteString(word)
			print(word)
		}
	}
	println()
	if err := stream.Err(); err != nil {
		return "", err
	}
	// 返回结果
	return s.String(), nil
}
