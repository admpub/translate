package libre

import (
	"context"
	"fmt"
	"strings"

	"github.com/admpub/translate"
	"github.com/webx-top/restyclient"
)

type libreRequest struct {
	Query string `json:"q"`
	CommonRequest
}

type libreBatchRequest struct {
	Query []string `json:"q"`
	CommonRequest
}

type libreDetectRequest struct {
	Query  string `json:"q"`
	Format string `json:"format,omitempty"`
	APIKey string `json:"api_key,omitempty"`
}

type CommonRequest struct {
	Source       string `json:"source"`
	Target       string `json:"target"`
	Format       string `json:"format"`
	Alternatives int    `json:"alternatives"`
	APIKey       string `json:"api_key"`
}

type libreResponse struct {
	Alternatives   []string `json:"alternatives"`
	TranslatedText string   `json:"translatedText"`
}

type libreBatchResponse struct {
	TranslatedText []string `json:"translatedText"`
}

type libreDetectResponse struct {
	Confidence float64 `json:"confidence"`
	Language   string  `json:"language"`
}

func fixLang(lang string) string {
	switch lang {
	case `zh-CN`:
		return `zh-Hans`
	case `zh-TW`, `zh-HK`:
		return `zh-Hant`
	default:
		return lang
	}
}

// libreTranslate translates text using the LibreTranslate API.
//
// API documentation: https://libretranslate.com/
//
//	APIConfig: {"apikey": "apikey", "endpoint": "https://libretranslate.com/translate"} or {"apikey": "apikey", "host": "libretranslate.com"}
func libreTranslate(ctx context.Context, cfg *translate.Config) (string, error) {
	comm := CommonRequest{
		Source: fixLang(cfg.From),
		Target: fixLang(cfg.To),
		Format: cfg.Format,
		APIKey: cfg.APIConfig[`apikey`],
	}
	data := libreRequest{
		Query:         cfg.Input,
		CommonRequest: comm,
	}
	var endpoint string
	if cfg.APIConfig[`endpoint`] != `` {
		endpoint = cfg.APIConfig[`endpoint`]
	} else {
		host := `libretranslate.com`
		if cfg.APIConfig[`host`] != `` {
			host = cfg.APIConfig[`host`]
		}
		endpoint = `https://` + host + `/translate`
	}
	r := &libreResponse{}
	req := restyclient.Retryable()
	req.SetContext(ctx)
	req.SetBody(data).SetResult(r).SetHeader(`Content-Type`, `application/json`).SetHeader(`Accept`, `application/json`)
	resp, e := req.Post(endpoint)
	if e != nil {
		return cfg.Input, e
	}
	if !resp.IsSuccess() {
		return cfg.Input, fmt.Errorf("[%v][%s] %s", resp.StatusCode(), resp.Status(), resp.Body())
	}
	return r.TranslatedText, nil
}

// DetectLanguage detects the language of the input text using the LibreTranslate API.
// It takes a context and translate.Config as input, and returns the detected language code or an error.
// The function constructs the API endpoint based on the configuration, sends a detection request,
// and processes the response to determine the language.
//
//	APIConfig: {"apikey": "apikey", "endpoint": "https://libretranslate.com/translate"} or {"apikey": "apikey", "host": "libretranslate.com"}
func DetectLanguage(ctx context.Context, cfg *translate.Config) (string, error) {
	var endpoint string
	if cfg.APIConfig[`endpoint`] != `` {
		endpoint = cfg.APIConfig[`endpoint`]
		endpoint = strings.TrimSuffix(endpoint, `/translate`)
		if !strings.HasSuffix(endpoint, `/detect`) {
			endpoint += `/detect`
		}
	} else {
		host := `libretranslate.com`
		if cfg.APIConfig[`host`] != `` {
			host = cfg.APIConfig[`host`]
		}
		endpoint = `https://` + host + `/detect`
	}
	data := &libreDetectRequest{
		Query:  cfg.Input,
		Format: cfg.Format,
		APIKey: cfg.APIConfig[`apikey`],
	}
	r := []*libreDetectResponse{}
	req := restyclient.Retryable()
	req.SetContext(ctx)
	req.SetBody(data).SetResult(&r).SetHeader(`Content-Type`, `application/json`).SetHeader(`Accept`, `application/json`)
	resp, e := req.Post(endpoint)
	if e != nil {
		return cfg.Input, e
	}
	if !resp.IsSuccess() {
		return cfg.Input, fmt.Errorf("[%v][%s] %s", resp.StatusCode(), resp.Status(), resp.Body())
	}
	if len(r) == 0 {
		return cfg.Input, fmt.Errorf("no response")
	}
	return r[0].Language, nil
}
