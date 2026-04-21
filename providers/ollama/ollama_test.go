package ollama

import (
	"context"
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
		From:   `zh-CN`,
		To:     `en`,
		Input:  `<h1>测试一下这个</h1>`,
		Format: `html`,
	})
	require.NoError(t, err)
	t.Log(result)
	require.Equal(t, `<h1>Test this out</h1>`, result)

	// - markdown -
	result, err = ollamaTranslate(context.Background(), &translate.Config{
		From: `zh-CN`,
		To:   `en`,
		Input: `# 测试一下这个
内容在这里`,
		Format: `markdown`,
	})
	require.NoError(t, err)
	t.Log(result)
	require.Equal(t, `# Let's test this out
Content goes here`, result)
}
