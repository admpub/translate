package ollama

import (
	"context"
	_ "embed"
	"strconv"
	"strings"

	"github.com/admpub/fasttemplate"
	"github.com/admpub/go-ollama"
	"github.com/admpub/translate"
	"github.com/coscms/forms/config"
)

func init() {
	translate.RegisterProvider(`ollama`, `Ollama`, ollamaTranslate, config.Element{
		Name:     `endpoint`,
		Type:     `url`,
		Label:    `Endpoint`,
		HelpText: `if a blank, use the local address: http://localhost:11434/api/generate`,
		Attributes: [][]string{
			{"placeholder", "http://localhost:11434/api/generate"},
		},
	}, config.Element{
		Name:  `apikey`,
		Type:  `text`,
		Label: `Token`,
	})
}

// $> curl -fsSL https://ollama.com/install.sh | sh
// $> ollama pull translategemma
// 翻译模型：https://ollama.com/library/translategemma
// 用Docker启动ollama：https://docs.ollama.com/docker

const (
	PromptTranslateGEMMA = `You are a professional {SOURCE_LANG} ({SOURCE_CODE}) to {TARGET_LANG} ({TARGET_CODE}) translator. Your goal is to accurately convey the meaning and nuances of the original {SOURCE_LANG} text while adhering to {TARGET_LANG} grammar, vocabulary, and cultural sensitivities.
Produce only the {TARGET_LANG} translation, without any additional explanations or commentary. Please translate the following {SOURCE_LANG} text into {TARGET_LANG}{EXTRA_INSTRUCTION}:


{TEXT}
`
)

var promptTemplate = fasttemplate.New(PromptTranslateGEMMA, `{`, `}`)

//go:embed SupportedLanguages.txt
var supportedLanguages string

var codeLanguages = map[string]string{}

func init() {
	for idx, row := range strings.Split(supportedLanguages, "\n") {
		if idx == 0 {
			continue
		}
		parts := strings.SplitN(row, "\t", 2)
		if len(parts) != 2 {
			continue
		}
		parts[0] = strings.TrimSpace(parts[0])
		parts[1] = strings.TrimSpace(parts[1])
		codeLanguages[parts[0]] = parts[1]
	}
	supportedLanguages = ``
}

func fixLangCode(code string) string {
	switch code {
	case `zh-CN`:
		return `zh-Hans`
	case `zh-HK`:
		return `zh-Hans-HK`
	default:
		return code
	}
}

func ollamaTranslate(ctx context.Context, cfg *translate.Config) (string, error) {
	cfg.To = fixLangCode(cfg.To)
	cfg.From = fixLangCode(cfg.From)
	sourceLang := codeLanguages[cfg.From]
	targetLang := codeLanguages[cfg.To]
	if len(cfg.Format) == 0 {
		cfg.Format = `text`
	}
	var extraInstruction string
	switch cfg.Format {
	case `html`:
		extraInstruction = ` and keep html tag unchanged`
	case `markdown`:
		extraInstruction = ` and keep markdown syntax`
	default:
	}
	prompt := promptTemplate.ExecuteString(map[string]interface{}{
		`SOURCE_LANG`:       sourceLang,
		`SOURCE_CODE`:       cfg.From,
		`TARGET_LANG`:       targetLang,
		`TARGET_CODE`:       cfg.To,
		`CONTENT_TYPE`:      cfg.Format,
		`EXTRA_INSTRUCTION`: extraInstruction,
		`TEXT`:              cfg.Input,
	})
	dsn := &ollama.DSN{
		URL:   cfg.GetAPIConfig(`url`, `endpoint`),
		Token: cfg.GetAPIConfig(`token`, `apikey`),
	}
	client := ollama.NewOpenWebUiClient(dsn)
	var result strings.Builder
	req := ollama.Request{
		Model:  "translategemma",
		Prompt: prompt,
		Options: &ollama.RequestOptions{
			Temperature: new(0.1),
			//NumContext: new(4096),
		},
		OnJson: func(res ollama.Response) error {
			if res.Response != nil {
				result.WriteString(*res.Response)
			}
			//ppnocolor.Println(res)
			return nil
		},
		/*//
		OnCodeBlock: func(cb []*ollama.CodeBlock) error {
			ppnocolor.Println(cb)
			return nil
		},
		//*/
	}
	cv, ok := cfg.APIConfig[`model`]
	if ok && len(cv) > 0 {
		req.Model = cv
	}
	cv, ok = cfg.APIConfig[`temperature`]
	if ok && len(cv) > 0 {
		if temperature, err := strconv.ParseFloat(cv, 64); err == nil {
			req.Options.Temperature = &temperature
		}
	}
	cv, ok = cfg.APIConfig[`numContext`]
	if ok && len(cv) > 0 {
		if numContext, err := strconv.Atoi(cv); err == nil {
			req.Options.NumContext = &numContext
		}
	}
	err := client.Query(req)
	return result.String(), err
}
