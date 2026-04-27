package singleton

import "sync"

type ConfigManager struct {
	mu     sync.RWMutex
	values map[string]string
}

var (
	instance *ConfigManager
	once     sync.Once
)

func GetInstance() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{
			values: make(map[string]string),
		}
	})
	return instance
}

func (c *ConfigManager) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
}

func (c *ConfigManager) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.values[key]
}
