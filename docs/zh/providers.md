# 提供者说明（中文）

本文件汇总了仓库中已实现的翻译提供者及其关键配置项，便于快速查阅。

- baidu
  - 描述：Baidu 翻译 API（支持传统 VIP 接口与 AI 翻译接口）。
  - 必需配置：`appid`, `secret`
  - 可选配置：`ai`（`true` 使用 AI 接口）
  - 参见实现：providers/baidu/translate_baidu.go

- tencent
  - 描述：腾讯云机器翻译（TMT）。
  - 必需配置：`appid`, `secret`
  - 参见实现：providers/tencent/translate_tencent.go

- google
  - 描述：使用 `gtranslate` 库调用 Google 翻译，支持自定义 host（例如 `google.cn`）。
  - 可选配置：`host`
  - 参见实现：providers/google/translate_google.go

- libre
  - 描述：LibreTranslate（可自托管）。
  - 必需/可选配置：`apikey`，可以直接指定 `endpoint`，或用 `host` + `scheme` 构造 URL。
  - 参见实现：providers/libre/translate_libre.go

- youdao
  - 描述：有道翻译（含 AI 接口变体）。
  - 必需配置：`appid`, `secret`
  - AI 变体可能支持 `prompt` 等额外配置。
  - 参见实现：providers/youdao/translate_youdao.go 和 translate_youdao_ai.go

- ollama
  - 描述：使用 Ollama 本地/私有部署的模型进行翻译（如 translategemma）。
  - 常用配置：`url`/`endpoint`, `token`/`apikey`, `model`, `temperature`, `numContext`
  - 参见实现：providers/ollama/ollama.go

如果需要添加新的 provider，只需实现 translate.Provider 的注册并在 `providers` 包或需要时通过 `init` 导入。
