package libre

import (
	"context"
	"os"
	"testing"

	"github.com/admpub/translate"
	"github.com/stretchr/testify/assert"
)

func TestLibre(t *testing.T) {
	text, err := libreTranslate(context.TODO(), translate.NewConfig(`测试`, `zh-Hans`, `en`, `text`).SetAPIConfig(`apikey`, os.Getenv(`LIBRE_APIKEY`)))
	assert.Equal(t, nil, err)
	assert.Equal(t, `test`, text)
}
