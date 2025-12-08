package google

import (
	"context"

	"github.com/admpub/translate"
	"github.com/bregydoc/gtranslate"
)

// googleTranslate translates the input text from source language to target language using Google Translate API with google.cn as the host.
//
//	APIConfig: {"host": "google.cn"}
func googleTranslate(ctx context.Context, cfg *translate.Config) (string, error) {
	host := cfg.APIConfig[`host`]
	if len(host) == 0 {
		host = `google.cn`
	}
	return gtranslate.TranslateWithParams(
		cfg.Input,
		gtranslate.TranslationParams{
			From:       cfg.From,
			To:         cfg.To,
			GoogleHost: host,
		},
	)
}
