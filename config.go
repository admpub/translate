package translate

type Config struct {
	Input     string
	From      string
	To        string
	Format    string
	APIConfig map[string]string
}

func (c *Config) SetAPIConfig(key, value string) *Config {
	c.APIConfig[key] = value
	return c
}

func NewConfig(input, from, to, format string) *Config {
	return &Config{
		Input:     input,
		From:      from,
		To:        to,
		Format:    format,
		APIConfig: map[string]string{},
	}
}
