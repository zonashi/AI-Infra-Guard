# AI Infrastructure Security Scan

## Feature Overview

The AI Infrastructure Security Scan module of AI Infra Guard is designed to detect known security vulnerabilities in web-based components within AI systems. By leveraging precise fingerprinting technology, this module can quickly identify security weaknesses in the AI infrastructure, helping security teams to promptly discover and remediate potential risks, thereby ensuring the secure and stable operation of AI systems.

## Core Features

- **Comprehensive Coverage**: Supports the identification of 36 mainstream AI component frameworks, covering 404 known vulnerabilities.
- **Flexible Deployment**: Supports single-target scanning, batch scanning, and one-click detection of local services.
- **Intelligent Matching**: A fingerprinting system based on YAML rules, ensuring high accuracy.
- **Extensibility**: Supports custom fingerprint rules and vulnerability templates to adapt to different environmental needs.

## Quick Start Guide

### Using the WebUI Interface

1. Click the "AI Infrastructure Scan" tab on the main interface.
2. Enter the URL or IP address to be scanned in the target input area.
   - Supports single-line or multi-line input (one target per line).
   - Supports importing target lists from a TXT file.
   - Entering an IP address will automatically scan all common open ports on that IP.
3. Click the "Start Scan" button, and the system will automatically perform the security check.

![image-20250717185311173](./assets/image-20250717185311173.png)

![image-20250717185509861](./assets/image-20250717185509861.png)

## Fingerprint and Vulnerability Database Management

### Built-in Fingerprint Database

AI Infra Guard comes with a rich built-in library of AI component fingerprints, which can be viewed and managed through the "Plugin Management" page:

1. Click "Plugin Management" in the bottom-left corner to enter the fingerprint management page.
2. On the fingerprint management page, you can view all built-in AI component fingerprint rules.
3. Operations such as searching, adding, and modifying fingerprints are supported.
By clicking on the plugin management page in the lower-left corner, you can see the built-in fingerprint and vulnerability databases of AIG.

![image-20250814173036377](./assets/image-20250814173036377.png)

In plugin management, you can search for fingerprints, corresponding vulnerabilities, add new ones, or modify existing ones. After modification, subsequent scans will use the updated fingerprint and vulnerability databases.

![image-20250717185223588](./assets/image-20250717185223588.png)

## Supported AI Components and Vulnerability Coverage

AI Infra Guard provides comprehensive security detection for critical components in AI infrastructure. The currently supported components and the number of vulnerabilities are as follows:

| Category                   | Component Name          | Vulnerability Count | Risk Level |
| -------------------------- | ----------------------- | ------------------- | ---------- |
| **Model Serving**          | gradio                  | 42                  | High       |
|                            | ollama                  | 7                   | Medium-High|
|                            | triton-inference-server | 7                   | Medium-High|
|                            | vllm                    | 4                   | Medium     |
|                            | xinference              | 0                   | Low        |
| **LLM App Frameworks**     | langchain               | 33                  | High       |
|                            | dify                    | 11                  | High       |
|                            | anythingllm             | 8                   | Medium-High|
|                            | open-webui              | 8                   | Medium-High|
|                            | ragflow                 | 2                   | Medium     |
|                            | qanything               | 2                   | Medium     |
| **Data Processing**        | clickhouse              | 22                  | High       |
|                            | feast                   | 0                   | Low        |
| **Visualization & UI**     | jupyter-server          | 13                  | Medium-High|
|                            | jupyterlab              | 6                   | Medium     |
|                            | jupyter-notebook        | 1                   | Low        |
|                            | tensorboard             | 0                   | Low        |
| **Workflow Orchestration** | kubeflow                | 4                   | Medium     |
|                            | ray                     | 4                   | Medium     |
| **Other AI Components**    | comfyui                 | 1                   | Low        |
|                            | comfy_mtb               | 1                   | Low        |
|                            | ComfyUI-Prompt-Preview  | 1                   | Low        |
|                            | ComfyUI-Custom-Scripts  | 1                   | Low        |
|                            | pyload-ng               | 18                  | Medium     |
|                            | kubepi                  | 5                   | Medium     |
|                            | llamafactory            | 1                   | Low        |
| **Total**                  |                         | **200+**            |            |

> **Note**: The vulnerability database is continuously updated. Regular scanning of high-risk components is recommended.

## Fingerprint Matching Rule Details

### Rule Structure

AI Infra Guard uses YAML format to define fingerprint matching rules, which mainly include the following parts:

```yaml
info:
  name: Component Name
  author: Rule Author
  severity: Information Level
  metadata:
    product: Product Name
    vendor: Vendor Name
http:
  - method: HTTP Request Method
    path: Request Path
    matchers:
      - Matching Conditions
```

### Example: Gradio Fingerprint Rule

```yaml
info:
  name: dify
  author: Tencent Zhuque Lab
  severity: info
  metadata:
    product: dify
    vendor: dify
http:
  - method: GET
    path: '/'
    matchers:
      - body="<title>Dify</title>" || icon="97378986"
version:
  - method: GET
    path: '/console/api/version'
    extractor:
      part: header
      group: 1
      regex: 'x-version:\s*(\d+\.\d+\.?\d+?)'
```

### Matcher Syntax Explanation

#### Match Locations

| Location | Description             | Example                                   |
| -------- | ----------------------- | ----------------------------------------- |
| `title`  | HTML page title         | `title="Gradio"`                          |
| `body`   | HTTP response body      | `body="gradio-config"`                    |
| `header` | HTTP response header    | `header="X-Gradio-Version: 3.34.0"`       |
| `icon`   | Website favicon hash    | `icon="d41d8cd98f00b204e9800998ecf8427e"` |

#### Logical Operators

| Operator | Description                               | Example                                                      |
| -------- | ----------------------------------------- | ------------------------------------------------------------ |
| `=`      | Fuzzy contains match (case-insensitive)   | `body="gradio"`                                              |
| `==`     | Exact equals match (case-sensitive)       | `header="Server: Gradio"`                                    |
| `!=`     | Not equals match                          | `header!="Server: Apache"`                                   |
| `~=`     | Regular expression match                  | `body~="Gradio v[0-9]+.[0-9]+.[0-9]+"`                       |
| `&&`     | Logical AND                               | `body="gradio" && header="X-Gradio-Version"`                 |
| `||`     | Logical OR                                | `body="gradio" || body="Gradio"`                             |
| `()`     | Grouping to change precedence             | `(body="gradio" || body="Gradio") && header="X-Gradio-Version"` |

## Best Practices

1.  **Regular Scanning**: It is recommended to perform a full scan of the AI infrastructure weekly to promptly discover new vulnerabilities.
2.  **Focus on High-Risk Components**: Components with a high number of vulnerabilities, such as gradio, langchain, and clickhouse, should be prioritized.
3.  **Extend with Custom Rules**: For enterprise-specific AI components, add custom fingerprint rules to enhance detection capabilities.
4.  **Integrate with CI/CD Pipelines**: Integrate security scanning into the continuous integration process of AI applications to achieve "shift-left" security.
5.  **Track Vulnerability Remediation**: Establish a tracking mechanism for vulnerabilities found during scans to ensure timely fixes.

By using the AI Infrastructure Security Scan module, you can effectively identify potential security risks in your AI systems, providing a strong guarantee for building a secure and reliable AI infrastructure.