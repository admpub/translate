package google

import (
	"testing"

	"github.com/admpub/translate"
	"github.com/stretchr/testify/assert"
)

func TestGoogle(t *testing.T) {
	text, err := googleTranslate(translate.NewConfig(`测试`, `zh-CN`, `en`, `text`))
	assert.Equal(t, nil, err)
	assert.Equal(t, `test`, text)
}
