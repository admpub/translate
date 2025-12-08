package google

import (
	"context"

	"github.com/admpub/translate"
	"github.com/bregydoc/gtranslate"
)

func googleTranslate(ctx context.Context, cfg *translate.Config) (string, error) {
	return gtranslate.TranslateWithParams(
		cfg.Input,
		gtranslate.TranslationParams{
			From:       cfg.From,
			To:         cfg.To,
			GoogleHost: `google.cn`,
		},
	)
}
