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
		Attributes: [][]string{
			{"required"},
		},
	}, config.Element{
		Name:  `secret`,
		Type:  `text`,
		Label: `Secret`,
		Attributes: [][]string{
			{"required"},
		},
	})
	translate.RegisterProvider(`youdaoAI`, `有道AI翻译`, youdaoAITranslate, config.Element{
		Name:  `appid`,
		Type:  `text`,
		Label: `App ID`,
		Attributes: [][]string{
			{"required"},
		},
	}, config.Element{
		Name:  `secret`,
		Type:  `text`,
		Label: `Secret`,
		Attributes: [][]string{
			{"required"},
		},
	}, config.Element{
		Name:  `prompt`,
		Type:  `text`,
		Label: `Prompt`,
	})
}
