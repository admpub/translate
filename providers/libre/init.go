package libre

import "github.com/admpub/translate"

func init() {
	translate.RegisterProvider(`libre`, `LibreTranslate`, libreTranslate)
}
