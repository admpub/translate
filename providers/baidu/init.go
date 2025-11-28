package baidu

import "github.com/admpub/translate"

func init() {
	translate.RegisterProvider(`baidu`, `百度翻译`, baiduTranslate)
}
