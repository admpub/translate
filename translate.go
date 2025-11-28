package translate

import "fmt"

var providers = map[string]*Provider{}

func RegisterProvider(name string, title string, translate func(cfg *Config) (string, error)) {
	providers[name] = &Provider{
		Name:      name,
		Title:     title,
		translate: translate,
	}
}

func GetProvider(name string) *Provider {
	return providers[name]
}

var ErrNotFound = fmt.Errorf("未找到翻译服务")

func Translate(name string, cfg *Config) (string, error) {
	provider := GetProvider(name)
	if provider == nil {
		return cfg.Input, fmt.Errorf("%w: %s", ErrNotFound, name)
	}
	return provider.Translate(cfg)
}
