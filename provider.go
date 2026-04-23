package translate

import (
	"context"

	"github.com/coscms/forms/config"
)

type Translator func(ctx context.Context, cfg *Config) (string, error)

type Provider struct {
	Name         string
	Title        string
	translate    Translator
	FormElements []config.Element
}

func (p *Provider) Translate(ctx context.Context, cfg *Config) (string, error) {
	return p.translate(ctx, cfg)
}
