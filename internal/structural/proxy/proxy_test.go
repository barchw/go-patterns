package proxy

import "testing"

func TestCachedWeatherServiceDelegatesOnMiss(t *testing.T) {
	real := NewRealWeatherService()
	var svc WeatherService = NewCachedWeatherService(real)

	result := svc.GetWeather("Berlin")
	if result == "" {
		t.Fatal("expected non-empty result")
	}
	if real.CallCount() != 1 {
		t.Fatalf("expected 1 call to real service, got %d", real.CallCount())
	}
}

func TestCachedWeatherServiceReturnsCachedOnHit(t *testing.T) {
	real := NewRealWeatherService()
	cached := NewCachedWeatherService(real)

	first := cached.GetWeather("Berlin")
	second := cached.GetWeather("Berlin")

	if first != second {
		t.Fatalf("cached result should match: %q != %q", first, second)
	}
	if real.CallCount() != 1 {
		t.Fatalf("expected 1 call to real service (second should be cached), got %d", real.CallCount())
	}
}

func TestCachedWeatherServiceDifferentCities(t *testing.T) {
	real := NewRealWeatherService()
	cached := NewCachedWeatherService(real)

	cached.GetWeather("Berlin")
	cached.GetWeather("Tokyo")
	cached.GetWeather("Berlin")
	cached.GetWeather("Tokyo")

	if real.CallCount() != 2 {
		t.Fatalf("expected 2 calls to real service (one per city), got %d", real.CallCount())
	}
}

func TestProxyImplementsInterface(t *testing.T) {
	real := NewRealWeatherService()
	var _ WeatherService = NewCachedWeatherService(real)
}
