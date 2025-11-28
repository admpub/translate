package translate

type Provider struct {
	Name      string
	Title     string
	translate func(cfg *Config) (string, error)
}

func (p *Provider) Translate(cfg *Config) (string, error) {
	return p.translate(cfg)
}
