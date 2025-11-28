package google

import "github.com/admpub/translate"

func init() {
	translate.RegisterProvider(`google`, `Google翻译`, googleTranslate)
}
