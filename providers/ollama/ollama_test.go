package ollama

import (
	"context"
	"testing"

	"github.com/admpub/translate"
	"github.com/stretchr/testify/require"
)

func TestOLLAMA(t *testing.T) {
	result, err := ollamaTranslate(context.Background(), &translate.Config{
		From:  `zh-CN`,
		To:    `en`,
		Input: `测试一下这个`,
	})
	require.NoError(t, err)
	t.Log(result)
	require.Equal(t, `Let's test this out.`, result)
}
