package libre

import (
	"github.com/admpub/translate"
	"github.com/coscms/forms/config"
)

func init() {
	translate.RegisterProvider(`libre`, `LibreTranslate`, libreTranslate, config.Element{
		Name:  `endpoint`,
		Type:  `url`,
		Label: `Endpoint`,
	}, config.Element{
		Name:  `apikey`,
		Type:  `text`,
		Label: `API Key`,
	})
}
