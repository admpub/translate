package tencent

import (
	"os"
	"testing"

	"github.com/admpub/translate"
	"github.com/stretchr/testify/assert"
)

func TestTencent(t *testing.T) {
	text, err := tencentTranslate(translate.NewConfig(`测试`, `zh-CN`, `en`, `text`).SetAPIConfig(`appid`, os.Getenv(`TENCENT_APPID`)).SetAPIConfig(`secret`, os.Getenv(`TENCENT_SECRET`)))
	assert.Equal(t, nil, err)
	assert.Equal(t, `test`, text)
}
