package libre

import (
	"context"
	"os"
	"testing"

	"github.com/admpub/translate"
	"github.com/stretchr/testify/assert"
)

func TestLibre(t *testing.T) {
	endpoint := `http://127.0.0.1:5000/translate`
	text, err := libreTranslate(context.TODO(), translate.NewConfig(`测试`, `zh-Hans`, `en`, `text`).SetAPIConfig(`apikey`, os.Getenv(`LIBRE_APIKEY`)).SetAPIConfig(`endpoint`, endpoint))
	assert.Equal(t, nil, err)
	assert.Equal(t, `Test`, text)
	text, err = DetectLanguage(context.TODO(), translate.NewConfig(`测试`, ``, ``, `text`).SetAPIConfig(`apikey`, os.Getenv(`LIBRE_APIKEY`)).SetAPIConfig(`endpoint`, endpoint))
	assert.Equal(t, nil, err)
	assert.Equal(t, `zh-Hans`, text)
}
