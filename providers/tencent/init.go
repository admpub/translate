package tencent

import (
	"github.com/admpub/translate"
	"github.com/coscms/forms/config"
)

func init() {
	translate.RegisterProvider(`tencent`, `腾讯翻译`, tencentTranslate, config.Element{
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
}
