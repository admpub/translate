package translate

import "sync"

// configPool is a sync.Pool for reusing Config instances.
var configPool = sync.Pool{
	New: func() interface{} {
		return &Config{inPool: true}
	},
}

// AcquireConfig returns a Config instance from the config pool.
func AcquireConfig() *Config {
	return configPool.Get().(*Config)
}

// ReleaseConfig releases the configuration back to the pool if it was allocated from the pool.
// If the config was not from the pool, this function does nothing.
// After releasing, the config is reset to its zero state.
func ReleaseConfig(c *Config) {
	c.Reset()
	configPool.Put(c)
}
