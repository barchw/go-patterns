package proxy

import "fmt"

type WeatherService interface {
	GetWeather(city string) string
}

type RealWeatherService struct {
	callCount int
}

func NewRealWeatherService() *RealWeatherService {
	return &RealWeatherService{}
}

func (r *RealWeatherService) GetWeather(city string) string {
	r.callCount++
	return fmt.Sprintf("Weather in %s: 22°C, Sunny", city)
}

func (r *RealWeatherService) CallCount() int {
	return r.callCount
}

type CachedWeatherService struct {
	realService *RealWeatherService
	cache       map[string]string
}

func NewCachedWeatherService(realService *RealWeatherService) *CachedWeatherService {
	return &CachedWeatherService{
		realService: realService,
		cache:       make(map[string]string),
	}
}

func (c *CachedWeatherService) GetWeather(city string) string {
	if result, ok := c.cache[city]; ok {
		return result
	}
	result := c.realService.GetWeather(city)
	c.cache[city] = result
	return result
}
