package ollama

import (
	"context"
	"strings"
	"testing"

	"github.com/admpub/translate"
	"github.com/stretchr/testify/require"
)

func TestOLLAMA(t *testing.T) {
	// - text -
	result, err := ollamaTranslate(context.Background(), &translate.Config{
		From:  `zh-CN`,
		To:    `en`,
		Input: `测试一下这个`,
	})
	require.NoError(t, err)
	t.Log(result)
	require.Equal(t, `Let's test this out.`, result)

	// - html -
	result, err = ollamaTranslate(context.Background(), &translate.Config{
		From: `zh-CN`,
		To:   `en`,
		Input: `<p>测试文字测试文字测试文字测试文字测试文字测试文字测试文字</p>
<p>测试文字测试文字测试文字测试文字测试文字测试文字</p>
<p>测试文字测试文字测试文字测试文字</p>
<p>测试文字测试文字测试文字测试文字</p>`,
		Format: `html`,
	})
	require.NoError(t, err)
	t.Log(result)
	require.Equal(t, `<p>Test text Test text Test text Test text Test text Test text Test text</p>
<p>Test text Test text Test text Test text Test text Test text</p>
<p>Test text Test text Test text Test text</p>
<p>Test text Test text Test text Test text</p>`, strings.TrimSpace(result))

	// - markdown -
	result, err = ollamaTranslate(context.Background(), &translate.Config{
		From: `zh-CN`,
		To:   `en`,
		Input: `# 测试一下这个
内容在这里
` + "```" + `
测试代码 测试代码 测试代码 测试代码 测试代码 测试代码 测试代码
` + "```" + `
`,
		Format: `markdown`,
	})
	require.NoError(t, err)
	t.Log(result)
	require.Equal(t, `# Let's test this
Content goes here
`+"```"+`
Test code Test code Test code Test code Test code Test code Test code
`+"```", result)
}
