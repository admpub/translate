package libre

import (
	"github.com/admpub/translate"
	"github.com/coscms/forms/config"
)

func init() {
	translate.RegisterProvider(`libre`, `LibreTranslate`, libreTranslate, config.Element{
		Name:     `endpoint`,
		Type:     `url`,
		Label:    `Endpoint`,
		HelpText: `if a blank, use the local address: http://127.0.0.1:5000/translate`,
		Attributes: [][]string{
			{"placeholder", "http://127.0.0.1:5000/translate"},
		},
	}, config.Element{
		Name:  `apikey`,
		Type:  `text`,
		Label: `API Key`,
	})
}
