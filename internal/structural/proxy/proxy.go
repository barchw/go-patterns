/*
Package proxy implements the Proxy design pattern.

What is it?
Proxy provides a substitute object that controls access to another object.
It implements the same interface as the target object and can add additional
behaviors (caching, logging, access control) before or after forwarding
the request to the real object. In this package, CachedWeatherService is
a caching proxy that prevents repeated calls to RealWeatherService.

When to use?
  - When you want to cache the results of expensive operations (caching proxy).
  - When you want to control access to an object (protection proxy).
  - When the real object is remote and you want to wrap the network communication.
  - When you want to delay the creation of an expensive object (lazy initialization).

When NOT to use?
  - When the extra layer of indirection provides no benefit.
  - When the target object is simple and cheap to use.
  - When you want to change the object's interface -- that's a job for the Adapter, not Proxy.

Tips and pitfalls:
  - Proxy vs. Decorator: Proxy controls access to an object, Decorator adds behaviors.
    In practice, the boundary is blurry.
  - Our example doesn't handle cache invalidation -- in production add TTL or a size limit.
  - The cache map is not thread-safe; in a concurrent program use sync.RWMutex.
  - A proxy should be transparent -- the client shouldn't know whether it's using the proxy or the real service.
*/
package proxy

import "fmt"

// WeatherService defines the common interface for the real weather service and the proxy.
type WeatherService interface {
	GetWeather(city string) string
}

// RealWeatherService simulates an expensive operation of querying an external weather API.
type RealWeatherService struct {
	callCount int
}

// NewRealWeatherService creates a new real weather service.
func NewRealWeatherService() *RealWeatherService {
	return &RealWeatherService{}
}

// GetWeather returns the weather for the given city and increments the call counter.
func (r *RealWeatherService) GetWeather(city string) string {
	r.callCount++
	return fmt.Sprintf("Weather in %s: 22°C, Sunny", city)
}

// CallCount returns the number of actual service calls (useful for verifying cache behavior).
func (r *RealWeatherService) CallCount() int {
	return r.callCount
}

// CachedWeatherService is a caching proxy that prevents repeated
// calls to the real service for the same city.
type CachedWeatherService struct {
	realService *RealWeatherService
	cache       map[string]string
}

// NewCachedWeatherService creates a new caching proxy wrapping the real weather service.
func NewCachedWeatherService(realService *RealWeatherService) *CachedWeatherService {
	return &CachedWeatherService{
		realService: realService,
		cache:       make(map[string]string),
	}
}

// GetWeather returns the weather from cache or queries the real service and stores the result.
func (c *CachedWeatherService) GetWeather(city string) string {
	if result, ok := c.cache[city]; ok {
		return result
	}
	result := c.realService.GetWeather(city)
	c.cache[city] = result
	return result
}
