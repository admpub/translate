package youdao

import (
	"github.com/admpub/translate"
	"github.com/coscms/forms/config"
)

func init() {
	translate.RegisterProvider(`youdao`, `有道文本翻译`, youdaoTranslate, config.Element{
		Name:  `appid`,
		Type:  `text`,
		Label: `App ID`,
	}, config.Element{
		Name:  `secret`,
		Type:  `text`,
		Label: `Secret`,
	})
	translate.RegisterProvider(`youdaoAI`, `有道AI翻译`, youdaoAITranslate, config.Element{
		Name:  `appid`,
		Type:  `text`,
		Label: `App ID`,
	}, config.Element{
		Name:  `secret`,
		Type:  `text`,
		Label: `Secret`,
	}, config.Element{
		Name:  `prompt`,
		Type:  `text`,
		Label: `Prompt`,
	})
}
