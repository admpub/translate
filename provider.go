package translate

import "context"

type Provider struct {
	Name      string
	Title     string
	translate func(ctx context.Context, cfg *Config) (string, error)
}

func (p *Provider) Translate(ctx context.Context, cfg *Config) (string, error) {
	return p.translate(ctx, cfg)
}
