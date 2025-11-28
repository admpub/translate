package youdao

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/admpub/translate"
	"github.com/webx-top/com"
	"github.com/webx-top/restyclient"
)

// documention: https://ai.youdao.com/DOCSIRMA/html/trans/api/dmxfy/index.html
type youdaoAIRequest struct {
	Input     string `json:"i"`
	From      string `json:"from"`
	To        string `json:"to"`
	Salt      string `json:"salt"`
	Sign      string `json:"sign"`
	Prompt    string `json:"prompt"`
	APPKey    string `json:"appKey"`
	CurrentTS string `json:"curtime"`

	// streamType 取值 	含义 	备注
	// increment 	译文文本按照“增量形式”返回 	默认按此方式返回
	// full 	译文文本按照“全量形式”返回
	// all 	同时返回“增量形式”、“全量形式”译文
	StreamType string `json:"streamType"`

	// handleOption取值 	含义
	// 0 	有道子曰翻译pro版本(14B)处理请求，通用pro翻译模型仅提供翻译功能，参数中的 prompt 仅对通用翻译模型(handleOption=0/3)生效
	// 3 	有道子曰翻译lite版本(1.5B)处理请求，lite翻译模型仅提供翻译功能，参数中的 prompt 仅对通用翻译模型(handleOption=0/3)生效
	HandleOption string `json:"handleOption"`
}

type youdaoAIResponse struct {
	Code       string               `json:"code"`
	Message    string               `json:"message"`
	Successful bool                 `json:"successful"`
	Data       youdaoAIResponseData `json:"data"`
	RequestId  string               `json:"requestId"`
}

type youdaoAIResponseData struct {
	TransIncrement string `json:"transIncre"`
	TransFull      string `json:"transFull"`
}

func youdaoAITranslate(cfg *translate.Config) (string, error) {
	time.Sleep(time.Second) // 接口频率限制：1次/秒
	data := youdaoAIRequest{
		Input:     cfg.Input,
		From:      fixLang(cfg.From),
		To:        fixLang(cfg.To),
		Salt:      com.RandomAlphanumeric(16),
		APPKey:    cfg.APIConfig[`appid`],
		Prompt:    cfg.APIConfig[`prompt`],
		CurrentTS: fmt.Sprint(time.Now().UTC().Unix()),
	}
	//input=q前10个字符 + q长度 + q后10个字符（当q长度大于20）或 input=q字符串（当q长度小于等于20）
	input := data.Input
	if len(data.Input) > 20 {
		input = data.Input[:10] + fmt.Sprint(len(data.Input)) + data.Input[len(data.Input)-10:]
	}
	data.Sign = com.Sha256(data.APPKey + input + data.Salt + data.CurrentTS + cfg.APIConfig[`secret`]) // 应用ID+input+salt+curtime+应用密钥
	endpoint := `https://openapi.youdao.com/proxy/http/llm-trans`
	req := restyclient.Classic().SetDoNotParseResponse(true)
	req.SetHeader(`Content-Type`, `application/json`).SetHeader(`Accept`, `text/event-stream`)
	//req.SetAuthToken(data.APPKey)
	resp, err := req.SetBody(data).Post(endpoint)
	if err != nil {
		return cfg.Input, err
	}
	defer resp.RawResponse.Body.Close()
	if !resp.IsSuccess() {
		return cfg.Input, fmt.Errorf("[%v][%s] %s", resp.StatusCode(), resp.Status(), resp.Body())
	}
	scanner := bufio.NewScanner(resp.RawResponse.Body)
	r := &youdaoAIResponse{}
	var translatedText string
	for scanner.Scan() {
		_res := scanner.Text()
		if _res == "" {
			continue
		}
		if after, found := strings.CutPrefix(_res, `data:`); found {
			err = json.Unmarshal([]byte(after), r)
			if err != nil {
				return cfg.Input, err
			}
			translatedText = r.Data.TransIncrement
		}
	}
	return translatedText, err
}
