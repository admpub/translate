package libre

import (
	"fmt"

	"github.com/admpub/translate"
	"github.com/webx-top/restyclient"
)

type libreRequest struct {
	Query        string `json:"q"`
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

func libreTranslate(cfg *translate.Config) (string, error) {
	data := libreRequest{
		Query:  cfg.Input,
		Source: fixLang(cfg.From),
		Target: fixLang(cfg.To),
		Format: cfg.Format,
		APIKey: cfg.APIConfig[`apikey`],
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
	req := restyclient.Classic()
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
