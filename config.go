package translate

// Config is the configuration for the translation service.
type Config struct {
	Input     string
	From      string
	To        string
	Format    string
	APIConfig map[string]string
	inPool    bool
}

// SetAPIConfig sets the API configuration value for the given key and returns the Config for chaining.
func (c *Config) SetAPIConfig(key, value string) *Config {
	c.APIConfig[key] = value
	return c
}

// Reset resets all configuration fields to their zero values and returns the Config pointer for method chaining.
func (c *Config) Reset() *Config {
	c.Input = ""
	c.From = ""
	c.To = ""
	c.Format = ""
	clear(c.APIConfig)
	return c
}

// Release returns the Config to the pool if it was allocated from one.
func (c *Config) Release() {
	if !c.inPool {
		return
	}
	ReleaseConfig(c)
}

// NewConfig creates a new Config instance with the specified input, source language, target language, and format.
// The returned Config includes an empty APIConfig map for additional provider-specific settings.
func NewConfig(input, from, to, format string) *Config {
	return &Config{
		Input:     input,
		From:      from,
		To:        to,
		Format:    format,
		APIConfig: map[string]string{},
	}
}
