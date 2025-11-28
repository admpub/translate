package youdao

import "github.com/admpub/translate"

func init() {
	translate.RegisterProvider(`youdao`, `有道文本翻译`, youdaoTranslate)
	translate.RegisterProvider(`youdaoAI`, `有道AI翻译`, youdaoAITranslate)
}
