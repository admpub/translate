package google

import (
	"github.com/admpub/translate"
	"github.com/coscms/forms/config"
)

func init() {
	translate.RegisterProvider(`google`, `Google翻译`, googleTranslate, config.Element{
		Name:     `host`,
		Type:     `text`,
		Label:    `Host`,
		HelpText: `Example: google.cn , Default: google.cn`,
	})
}
