package translate

import (
	"context"
	"fmt"
	"sort"

	"github.com/coscms/forms/config"
)

var providers = map[string]*Provider{}

func RegisterProvider(name string, title string, translate Translator, formElements ...config.Element) {
	providers[name] = &Provider{
		Name:         name,
		Title:        title,
		translate:    translate,
		FormElements: formElements,
	}
}

func GetProvider(name string) *Provider {
	return providers[name]
}

var ErrNotFound = fmt.Errorf("未找到翻译服务")

func Translate(ctx context.Context, name string, cfg *Config) (string, error) {
	provider := GetProvider(name)
	if provider == nil {
		return cfg.Input, fmt.Errorf("%w: %s", ErrNotFound, name)
	}
	return provider.Translate(ctx, cfg)
}

func FormElements(ctx context.Context, name string) []config.Element {
	provider := GetProvider(name)
	if provider == nil {
		return nil
	}
	return provider.FormElements
}

func ListAll() []Provider {
	names := make([]string, 0, len(providers))
	for name := range providers {
		names = append(names, name)
	}
	sort.Strings(names)
	results := make([]Provider, len(names))
	for index, name := range names {
		results[index] = *providers[name]
	}
	return results
}
