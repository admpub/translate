# 使用说明与示例（中文）

## 基本用法

示例：使用 Baidu 翻译

```go
package main

import (
    "context"
    "fmt"
    "github.com/admpub/translate"
    _ "github.com/admpub/translate/providers"
)

func main() {
    cfg := translate.NewConfig("中国", "zh-CN", "en", "text")
    cfg.SetAPIConfig("appid", "your-appid")
    cfg.SetAPIConfig("secret", "your-secret")
    res, err := translate.Translate(context.Background(), "baidu", cfg)
    if err != nil {
        fmt.Println("翻译错误：", err)
        return
    }
    fmt.Println("翻译结果：", res)
}
```

## 进阶提示

- 链式设置：`SetAPIConfig` 返回 `*Config`，可连续调用。
- 如果使用私有部署（如 LibreTranslate 或 Ollama），请在 `APIConfig` 中设置 `endpoint`/`url` 与凭证（`apikey`/`token`）。
- 某些 provider 会自动调整语言代码（例如 `zh-CN` -> `zh-Hans`），无须额外处理。
