# 🛡️ A.I.G (AI-Infra-Guard)

<div align="center">

[English](./README_EN.md) ｜ **中文** ｜ [帮助文档](https://tencent.github.io/AI-Infra-Guard/)



  [![GitHub Stars](https://img.shields.io/github/stars/Tencent/AI-Infra-Guard?style=for-the-badge&logo=github&color=gold)](https://github.com/Tencent/AI-Infra-Guard)
  [![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)](LICENSE)
  [![Release](https://img.shields.io/github/v/release/Tencent/AI-Infra-Guard?style=for-the-badge&color=green)](https://github.com/Tencent/AI-Infra-Guard/releases)

  <h3>🚀 腾讯朱雀实验室出品的一站式 AI 红队安全测试平台</h3>

  
</div>

**A.I.G (AI-Infra-Guard)** 集成AI基础设施漏洞扫描、MCP Server风险检测与大模型安全体检等能力，旨在为用户提供最全面、智能与易用的AI安全风险自查解决方案。


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
      <p>精准识别 <b>30+</b> AI 框架组件<br/>覆盖 <b>近400个</b> 已知 CVE 漏洞<br/>如 Ollama、ComfyUI、vLLM 等</p>
    </td>
    <td align="center" width="33%">
      <h3>🤖 MCP Server 安全检测</h3>
      <p>AI Agent 驱动的智能分析<br/>检测 <b>9 大类</b>MCP安全风险<br/>支持MCP <b>源代码/远程URL</b> 等多种扫描方式</p>
    </td>
    <td align="center" width="33%">
      <h3>⚡ 大模型安全体检</h3>
      <p>一键评估多个模型Prompt安全性<br/>内置多个主流&独家评测集<br/>快速获取模型安全体检报告</p>
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

### 插件管理
<div>
<div style="display: flex;">
  <img src="img/plugin-1.png" alt="大模型安全体检" style="width: 50%; margin-right: 5px;">
  <img src="img/plugin-2.png" alt="AI基础设施扫描" style="width: 50%; margin-left: 5px;">
</div>
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

然后，并安装Docker相关环境，您可以选择以下任一方式部署A.I.G服务：

**1. 使用预构建镜像运行 (推荐)**
```bash
# 此方法会从 Docker Hub 拉取预构建的镜像，启动速度更快
docker-compose -f docker-compose.images.yml up -d
```
**2. 从源码构建并运行**
```bash
# 此方法会使用本地代码构建 Docker 镜像并启动服务
docker-compose up -d
```

服务启动后，即可访问A.I.G的 Web 界面：
`http://localhost:8088`

## 📝 贡献指南

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
- ⭐ [改进文档](https://github.com/Tencent/AI-Infra-Guard/pulls)

### 🙏 致谢

感谢以下安全团队的专业共建:

<div style="display: flex; justify-content: center; align-items: center; gap: 2%;">
  <img src="img/keen_lab_logo.svg" alt="Keen Lab" style="width: 30%; height: auto;">
  <img src="img/wechat_security.png" alt="WeChat Security" style="width: 30%; height: auto;">
  <img src="img/fit_sec_logo.png" alt="Fit Security" style="width: 30%; height: auto;">
</div>
</br>

感谢为 A.I.G 项目贡献代码的开发者：

<a href="https://github.com/Tencent/AI-Infra-Guard/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Tencent/AI-Infra-Guard" />
</a>



## 💬 加入社区

<div align="center">
  <img src="img/wechatgroup.png" alt="微信群" width="200">
  <p><b>微信扫码加入 A.I.G 交流群</b></p>
  <p>共同探讨 AI 安全最佳实践</p>
</div>

## 📄 开源协议

本项目基于 **MIT 许可证** 开源。详细信息请查阅 [License.txt](./License.txt) 文件。

---

<div align="center">

  [![Star History Chart](https://api.star-history.com/svg?repos=Tencent/AI-Infra-Guard&type=Date)](https://star-history.com/#Tencent/AI-Infra-Guard&Date)

</div>
