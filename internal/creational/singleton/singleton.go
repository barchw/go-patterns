/*
Package singleton implements the Singleton pattern.

What is it?
Singleton guarantees that a given type has exactly one instance and provides
a global point of access to it. In Go it is idiomatically implemented using
sync.Once, which ensures thread safety and lazy initialization.

When to use?
  - When the entire application needs exactly one instance (e.g., configuration, cache, connection pool).
  - When object creation is expensive and you want to initialize it lazily.
  - When multiple goroutines may simultaneously try to access a shared resource.

When NOT to use?
  - When the Singleton masks hidden dependencies -- it makes testing and data-flow tracking harder.
  - When you need different configurations in different parts of the application.
  - When you can use dependency injection instead.

Tips and pitfalls:
  - sync.Once is the Go idiom for Singleton -- do not implement it via double-checked locking or init().
  - A package-level variable survives between tests -- consider resetForTesting() or dependency injection.
  - Using RWMutex is crucial when reads greatly outnumber writes.
  - In modern Go applications, prefer creating the instance in main() and passing it via DI.
*/
package singleton

import "sync"

// ConfigManager is a thread-safe key-value configuration manager (Singleton).
type ConfigManager struct {
	mu     sync.RWMutex
	values map[string]string
}

var (
	instance *ConfigManager
	once     sync.Once
)

// GetInstance returns the sole ConfigManager instance, creating it lazily on the first call.
func GetInstance() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{
			values: make(map[string]string),
		}
	})
	return instance
}

// Set stores a value for the given configuration key (thread-safe).
func (c *ConfigManager) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.values[key] = value
}

// Get returns the value for the given configuration key (thread-safe, uses RLock).
func (c *ConfigManager) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.values[key]
}
