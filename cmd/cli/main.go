package main

import (
	"github.com/Tencent/AI-Infra-Guard/common/runner"
	"github.com/Tencent/AI-Infra-Guard/internal/gologger"
	"github.com/Tencent/AI-Infra-Guard/internal/options"
)

func main() {
	parseOptions := options.ParseOptions() // 解析命令行参数
	r, err := runner.New(parseOptions)     // 创建runner
	if err != nil {
		// 打印错误
		gologger.Fatalf("Could not create runner: %s\n", err)
	}
	defer r.Close() // 关闭runner
	gologger.Logger.SetLevel(gologger.InfoLevel)
	if parseOptions.ListFpsTemplate || parseOptions.ListVulTemplate {
		// 打印指纹和漏洞列表
		r.ShowFpAndVulList(parseOptions.ListFpsTemplate, parseOptions.ListVulTemplate)
		return
	}
	r.RunEnumeration() // 执行枚举
}
