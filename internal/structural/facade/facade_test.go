package facade

import (
	"strings"
	"testing"
)

func TestMediaConverterFacade(t *testing.T) {
	converter := NewMediaConverter()
	result := converter.Convert("movie.avi", "mp4")

	if !strings.Contains(result, "movie.avi") {
		t.Fatalf("expected result to reference input file, got %s", result)
	}
	if !strings.HasSuffix(result, ".mp4") {
		t.Fatalf("expected result to end with .mp4, got %s", result)
	}
	if !strings.Contains(result, "processed_video") {
		t.Fatalf("expected result to show video processing, got %s", result)
	}
	if !strings.Contains(result, "audio_track") {
		t.Fatalf("expected result to show audio extraction, got %s", result)
	}

	expected := "processed_video(movie.avi)+audio_track(movie.avi).mp4"
	if result != expected {
		t.Fatalf("got %q, want %q", result, expected)
	}
}
