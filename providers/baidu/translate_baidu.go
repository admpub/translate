package baidu

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/admpub/translate"
	"github.com/webx-top/com"
	"github.com/webx-top/restyclient"
)

/*
{
    "from": "zh",
    "to": "en",
    "trans_result": [
        {
            "src": "中国",
            "dst": "China"
        }
    ]
}
*/

type baiduTransResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}
type baiduResponse struct {
	From    string              `json:"from"`
	To      string              `json:"to"`
	Results []*baiduTransResult `json:"trans_result"`
}

// documention: https://fanyi-api.baidu.com/product/113
func baiduTranslate(cfg *translate.Config) (string, error) {
	time.Sleep(time.Second) // 接口频率限制：1次/秒
	values := url.Values{
		`q`:     []string{cfg.Input},
		`from`:  []string{strings.SplitN(cfg.From, `-`, 2)[0]},
		`to`:    []string{strings.SplitN(cfg.To, `-`, 2)[0]},
		`appid`: []string{cfg.APIConfig[`appid`]},
		`salt`:  []string{com.RandomAlphanumeric(16)},
	}
	sign := com.Md5(cfg.APIConfig[`appid`] + values.Get(`q`) + values.Get(`salt`) + cfg.APIConfig[`secret`]) //  appid+q+salt+密钥
	values.Add(`sign`, sign)
	aiTranslate, _ := strconv.ParseBool(cfg.APIConfig[`ai`])
	var endpoint string
	if aiTranslate {
		endpoint = `https://fanyi-api.baidu.com/ait/api/aiTextTranslate`
	} else {
		endpoint = `https://fanyi-api.baidu.com/api/trans/vip/translate`
	}
	req := restyclient.Classic()
	resp, e := req.SetFormDataFromValues(values).Post(endpoint)
	if e != nil {
		return cfg.Input, e
	}
	if !resp.IsSuccess() {
		return cfg.Input, fmt.Errorf("[%v][%s] %s", resp.StatusCode(), resp.Status(), resp.Body())
	}
	r := &baiduResponse{}
	err := json.Unmarshal(resp.Body(), r)
	if err != nil {
		return cfg.Input, err
	}
	//com.Dump(r)
	for _, v := range r.Results {
		return v.Dst, nil
	}
	return cfg.Input, nil
}
