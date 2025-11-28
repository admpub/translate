package tencent

import "github.com/admpub/translate"

func init() {
	translate.RegisterProvider(`tencent`, `腾讯翻译`, tencentTranslate)
}
