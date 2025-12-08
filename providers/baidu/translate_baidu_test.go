package baidu

import (
	"context"
	"os"
	"testing"

	"github.com/admpub/translate"
	"github.com/stretchr/testify/assert"
)

func TestBaidu(t *testing.T) {
	text, err := baiduTranslate(context.TODO(), translate.NewConfig(`测试`, `zh-CN`, `en`, `text`).SetAPIConfig(`appid`, os.Getenv(`BAIDU_APPID`)).SetAPIConfig(`secret`, os.Getenv(`BAIDU_SECRET`)))
	assert.Equal(t, nil, err)
	assert.Equal(t, `test`, text)
}
