package tencent

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/admpub/translate"
	"github.com/webx-top/restyclient"
)

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacsha256(s, key string) string {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(s))
	return string(hashed.Sum(nil))
}

type tencentRequest struct {
	SourceTextList []string
	Source         string // zh
	Target         string // en
	ProjectId      int
}

type tencentResponse struct {
	Response tencentResponseData
}

type tencentResponseData struct {
	RequestId      string
	Source         string
	Target         string
	TargetTextList []string
	Error          *tencentResponseError
}

type tencentResponseError struct {
	Code    string
	Message string
}

const tencentMaxBytes = 6000

// documention: https://cloud.tencent.com/document/product/551/40566
func tencentTranslate(ctx context.Context, cfg *translate.Config) (string, error) {
	url := `https://tmt.tencentcloudapi.com`
	req := restyclient.Retryable()
	data := tencentRequest{
		SourceTextList: []string{cfg.Input},
		Source:         strings.SplitN(cfg.From, `-`, 2)[0],
		Target:         strings.SplitN(cfg.To, `-`, 2)[0],
	}
	body, _ := json.Marshal(data)
	secretId := cfg.APIConfig[`appid`]
	secretKey := cfg.APIConfig[`secret`]
	r := &tencentResponse{}
	req.SetContext(ctx)
	req.SetBody(data).SetResult(r).SetHeaders(makeTencentSign(secretId, secretKey, string(body)))
	resp, e := req.Post(url)
	if e != nil {
		return cfg.Input, e
	}
	if !resp.IsSuccess() {
		return cfg.Input, fmt.Errorf("[%v][%s] %s", resp.StatusCode(), resp.Status(), resp.Body())
	}
	if r.Response.Error != nil {
		return cfg.Input, fmt.Errorf("[%s] %s", r.Response.Error.Code, r.Response.Error.Message)
	}
	//fmt.Println(resp.String())
	//com.Dump(r)
	for _, v := range r.Response.TargetTextList {
		return v, nil
	}
	return cfg.Input, nil
}

// documention: https://cloud.tencent.com/document/api/213/30654#Golang
func makeTencentSign(secretId, secretKey, payload string) map[string]string {
	host := "tmt.tencentcloudapi.com"
	algorithm := "TC3-HMAC-SHA256"
	service := "tmt"
	version := "2018-03-21"
	action := "TextTranslateBatch"
	region := "ap-guangzhou"
	var timestamp int64 = time.Now().Unix()

	// step 1: build canonical request string
	httpRequestMethod := "POST"
	canonicalURI := "/"
	canonicalQueryString := ""
	canonicalHeaders := fmt.Sprintf("content-type:%s\nhost:%s\nx-tc-action:%s\n",
		"application/json; charset=utf-8", host, strings.ToLower(action))
	signedHeaders := "content-type;host;x-tc-action"
	//payload := `{"Limit": 1, "Filters": [{"Values": ["\u672a\u547d\u540d"], "Name": "instance-name"}]}`
	hashedRequestPayload := sha256hex(payload)
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload)
	//fmt.Println(canonicalRequest)

	// step 2: build string to sign
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	credentialScope := fmt.Sprintf("%s/%s/tc3_request", date, service)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	string2sign := fmt.Sprintf("%s\n%d\n%s\n%s",
		algorithm,
		timestamp,
		credentialScope,
		hashedCanonicalRequest)
	//fmt.Println(string2sign)

	// step 3: sign string
	secretDate := hmacsha256(date, "TC3"+secretKey)
	secretService := hmacsha256(service, secretDate)
	secretSigning := hmacsha256("tc3_request", secretService)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, secretSigning)))
	//fmt.Println(signature)

	// step 4: build authorization
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		secretId,
		credentialScope,
		signedHeaders,
		signature)
	//fmt.Println(authorization)
	/*//
		curl := fmt.Sprintf(`curl -X POST https://%s\
		 -H "Authorization: %s"\
		 -H "Content-Type: application/json; charset=utf-8"\
		 -H "Host: %s" -H "X-TC-Action: %s"\
		 -H "X-TC-Timestamp: %d"\
		 -H "X-TC-Version: %s"\
		 -H "X-TC-Region: %s"\
		 -d '%s'`, host, authorization, host, action, timestamp, version, region, payload)
		fmt.Println(curl)
	//*/
	return map[string]string{
		`Content-Type`:   `application/json; charset=utf-8`,
		`Authorization`:  authorization,
		`X-TC-Action`:    action,
		`X-TC-Timestamp`: strconv.FormatInt(timestamp, 10),
		`X-TC-Version`:   version,
		`X-TC-Region`:    region,
	}
}
