package baidu

import (
	"github.com/admpub/translate"
	"github.com/coscms/forms/config"
)

func init() {
	translate.RegisterProvider(`baidu`, `百度翻译`, baiduTranslate, config.Element{
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
