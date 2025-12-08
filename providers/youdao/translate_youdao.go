package youdao

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/admpub/translate"
	"github.com/webx-top/com"
	"github.com/webx-top/restyclient"
)

func fixLang(lang string) string {
	switch lang {
	case `zh-CN`:
		return `zh-CHS`
	case `zh-TW`, `zh-HK`:
		return `zh-CHT`
	default:
		return lang
	}
}

type youdaoRequest struct {
	Query     string `json:"q"`
	From      string `json:"from"`
	To        string `json:"to"`
	Salt      string `json:"salt"`
	Sign      string `json:"sign"`
	SignType  string `json:"signType"`
	APPKey    string `json:"appKey"`
	CurrentTS string `json:"curtime"`
}

func (y *youdaoRequest) ToValues() url.Values {
	return url.Values{
		`q`:        []string{y.Query},
		`from`:     []string{y.From},
		`to`:       []string{y.To},
		`salt`:     []string{y.Salt},
		`sign`:     []string{y.Sign},
		`signType`: []string{y.SignType},
		`appKey`:   []string{y.APPKey},
		`curtime`:  []string{y.CurrentTS},
	}
}

type youdaoResponse struct {
	ErrorCode   string   `json:"errorCode"`
	Translation []string `json:"translation"`
}

// youdaoTranslate performs translation using Youdao API.
//
// API documentation: https://ai.youdao.com/DOCSIRMA/html/trans/api/wbfy/index.html
//
//	APIConfig: {"appid": "appid", "secret": "secret"}
func youdaoTranslate(ctx context.Context, cfg *translate.Config) (string, error) {
	data := youdaoRequest{
		Query:     cfg.Input,
		From:      fixLang(cfg.From),
		To:        fixLang(cfg.To),
		Salt:      com.RandomAlphanumeric(16),
		APPKey:    cfg.APIConfig[`appid`],
		SignType:  `v3`,
		CurrentTS: fmt.Sprint(time.Now().UTC().Unix()),
	}
	//input=q前10个字符 + q长度 + q后10个字符（当q长度大于20）或 input=q字符串（当q长度小于等于20）
	input := data.Query
	if len(data.Query) > 20 {
		input = data.Query[:10] + fmt.Sprint(len(data.Query)) + data.Query[len(data.Query)-10:]
	}
	data.Sign = com.Sha256(data.APPKey + input + data.Salt + data.CurrentTS + cfg.APIConfig[`secret`]) // 应用ID+input+salt+curtime+应用密钥
	endpoint := `https://openapi.youdao.com/api`
	r := &youdaoResponse{}
	req := restyclient.Retryable()
	req.SetContext(ctx)
	req.SetHeader(`Accept`, `application/json`)
	//req.SetAuthToken(data.APPKey)
	//req.SetBody(data)
	req.SetFormDataFromValues(data.ToValues())
	resp, e := req.SetResult(r).Post(endpoint)
	if e != nil {
		return cfg.Input, e
	}
	if !resp.IsSuccess() {
		return cfg.Input, fmt.Errorf("[%v][%s] %s", resp.StatusCode(), resp.Status(), resp.Body())
	}
	for _, v := range r.Translation {
		return v, nil
	}
	return cfg.Input, nil
}
