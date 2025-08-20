# 🛡️ A.I.G (AI-Infra-Guard)

<div align="center">

  [![GitHub Stars](https://img.shields.io/github/stars/Tencent/AI-Infra-Guard?style=for-the-badge&logo=github&color=gold)](https://github.com/Tencent/AI-Infra-Guard)
  [![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)](LICENSE)
  [![Release](https://img.shields.io/github/v/release/Tencent/AI-Infra-Guard?style=for-the-badge&color=green)](https://github.com/Tencent/AI-Infra-Guard/releases)

  <h3>🚀 腾讯朱雀实验室出品的一站式 AI 安全红队测试平台</h3>


  <strong>⭐ 如果觉得有用，请给我们一个 Star 支持开源！</strong>
  
</div>

**A.I.G (AI-Infra-Guard)** 旨在为AI安全提供一站式的安全解决方案。A.I.G集成常见AI安全，通过全面、智能、易用的AI安全检测工具，帮助开发者和企业发现并修复 AI 系统中的安全漏洞，确保 AI 应用的安全、可靠、可控。

[English](./README_EN.md) | **简体中文** ｜ [帮助文档](https://tencent.github.io/AI-Infra-Guard/)

## 目录
- [✨ 核心功能](#-核心功能)
- [🖼️ 功能展示](#-功能展示)
- [🚀 快速开始](#-快速开始)
- [🙏 贡献指南](#-贡献指南)
- [💬 加入社区](#-加入社区)
- [📄 开源协议](#-开源协议)

## ✨ 核心功能

<table>
  <tr>
    <td align="center" width="33%">
      <h3>🔍 AI 基础设施扫描</h3>
      <p>精准识别 <b>28+</b> AI 框架组件<br/>覆盖 <b>200+</b> 已知 CVE 漏洞<br/>支持 Gradio、Ollama、Dify、ComfyUI 等</p>
    </td>
    <td align="center" width="33%">
      <h3>🤖 MCP Server 安全检测</h3>
      <p>AI Agent 驱动的智能分析<br/>检测 <b>9 大类</b>安全风险<br/>支持 <b>源码/远程</b> 等多种扫描方式</p>
    </td>
    <td align="center" width="33%">
      <h3>⚡ 大模型安全体检</h3>
      <p>基于prompt对话场景评测<br/>内置多种prompt评测方法<br/>一键生成专业安全报告</p>
    </td>
  </tr>
</table>

## 🖼️ 功能展示

### AI 基础设施漏洞扫描
<div align="center">
  <img src="img/ai-infra-scan.png" alt="AI基础设施扫描" width="600">
  <p><i>一键扫描发现 AI 组件安全漏洞</i></p>
</div>

### MCP Server 安全检测
<div align="center">
  <img src="img/mcp-server.png" alt="MCP安全检测" width="600">
  <p><i>智能分析 MCP Server 安全风险</i></p>
</div>

### 大模型安全体检
<div align="center">
  <img src="img/model-security.png" alt="大模型安全体检" width="600">
  <p><i>全面评估大模型安全性</i></p>
</div>

## 🚀 快速开始
### Docker 一键部署

**系统要求**

- Docker 20.10 或更高版本
- 至少 4GB 可用内存
- 至少 10GB 可用磁盘空间

首先，克隆本项目：
```bash
git clone https://github.com/Tencent/AI-Infra-Guard.git
cd AI-Infra-Guard
```

然后，您可以选择以下任一方式启动服务：

**1. 从源码构建并运行**
```bash
# 此方法会使用本地代码构建 Docker 镜像并启动服务
docker-compose up -d
```

**2. 使用预构建镜像运行 (推荐)**
```bash
# 此方法会从 Docker Hub 拉取预构建的镜像，启动速度更快
docker-compose -f docker-compose.images.yml up -d
```

服务启动后，即可访问 Web 界面：
`http://localhost:8088`

## 🙏 贡献指南

A.I.G 的核心能力之一就是其丰富且可快速配置的插件系统。我们欢迎社区贡献高质量的插件和功能。

### 贡献插件规则
1.  **指纹规则**: 在 `data/fingerprints/` 目录下添加新的 YAML 指纹文件
2.  **漏洞规则**: 在 `data/vuln/` 目录下添加新的漏洞检测规则
3.  **MCP 插件**: 在 `data/mcp/` 目录下添加新的 MCP 安全检测规则
4.  **模型评测集**: 在 `data/eval` 目录下添加新的模型评测集

请参考现有规则格式，创建新文件并通过 Pull Request 提交。

### 其他贡献方式
- 🐛 [报告 Bug](https://github.com/Tencent/AI-Infra-Guard/issues)
- 💡 [提出新功能](https://github.com/Tencent/AI-Infra-Guard/issues)
- 📝 [改进文档](https://github.com/Tencent/AI-Infra-Guard/pulls)
- ⭐ 给项目点个 Star

### 🙏 致谢

感谢以下参与A.I.G共建团队或项目:
- xx

感谢所有为 A.I.G 项目做出贡献的开发者们：

<a href="https://github.com/Tencent/AI-Infra-Guard/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Tencent/AI-Infra-Guard" />
</a>



## 💬 加入社区

<div align="center">
  <img src="img/wechatgroup.png" alt="微信群" width="200">
  <p><b>扫码加入 A.I.G 技术交流群</b></p>
  <p>与开发者一起讨论 AI 安全最佳实践</p>
</div>

## 📄 开源协议

本项目基于 [MIT License](./LICENSE) 开源。


<div align="center">

  [![Star History Chart](https://api.star-history.com/svg?repos=Tencent/AI-Infra-Guard&type=Date)](https://star-history.com/#Tencent/AI-Infra-Guard&Date)

</div>
