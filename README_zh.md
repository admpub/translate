# translate（中文文档）

`translate` 是一个轻量的 Go 翻译封装库，集成了多个翻译提供者（Baidu、Tencent、LibreTranslate、Youdao、Ollama、Google 等），提供统一的 `Config` 配置和 `Translate` 调用接口，方便在 Go 应用中调用不同翻译服务。

## 快速开始

1. 在你的 Go 项目中引入并导入 providers（init 注册）：

   import (
       "context"
       "github.com/admpub/translate"
       _ "github.com/admpub/translate/providers"
   )

2. 示例代码：

```go
cfg := translate.NewConfig("你好世界", "zh-CN", "en", "text")
cfg.SetAPIConfig("appid", "你的appid")
cfg.SetAPIConfig("secret", "你的secret")
res, err := translate.Translate(context.Background(), "baidu", cfg)
if err != nil {
    // 处理错误
}
fmt.Println(res)
```

注意：大多数 provider 需要你在 `cfg.APIConfig` 中设置 provider 指定的密钥或 endpoint。

## 核心类型与方法

- `translate.Config`：主配置结构体，字段包括 `Input`, `From`, `To`, `Format`, `APIConfig` 等。
- `translate.NewConfig(input, from, to, format)`：创建新的配置对象。
- `(*Config).SetAPIConfig(key, value)`：设置 API 专用配置，支持链式调用。
- `translate.Translate(ctx, providerName, cfg)`：执行翻译，`providerName` 为提供者名称（如 `baidu`、`tencent`、`libre`、`youdao`、`ollama`）。

## 支持的提供者与配置项

- `baidu`：APIConfig: `appid`, `secret`, `ai`（可选，`true` 使用 AI 接口）
- `tencent`：APIConfig: `appid`, `secret`
- `google`：APIConfig: `host`（可选，默认为 `google.cn`）
- `libre`：APIConfig: `apikey`, `endpoint` 或 `host` + `scheme`
- `youdao`：APIConfig: `appid`, `secret`, （Youdao AI 变体可能支持 `prompt`）
- `ollama`：APIConfig: `url`/`endpoint`, `token`/`apikey`, 可选 `model`, `temperature`, `numContext`

更多细节参见 `providers` 目录中每个 provider 的实现。

## 语言代码与格式

请用常见的语言代码（如 `zh-CN`, `en`, `ja` 等）。部分 provider 对某些中文变体会进行内部修正（例如 `zh-CN` -> `zh-Hans`）。

## 用例与建议

- 若仅作简单翻译，可以直接调用 `Translate` 并传入相应 provider 的 API 配置。
- 对于自托管或私有服务（如 Ollama、LibreTranslate 私有部署），请在 `APIConfig` 中设置 `endpoint`/`url` 并配置 `token` 或 `apikey`。

---

仓库源码入口：参见 `providers` 目录以了解各 provider 的更多实现细节。
