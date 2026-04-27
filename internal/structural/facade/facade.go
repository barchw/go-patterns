/*
Package facade implements the Facade design pattern.

What is it?
Facade provides a simplified interface to a complex subsystem. Instead of forcing
the client to interact with many classes, the facade offers a single convenient method
that coordinates all the work. In this package, MediaConverter is a facade hiding
the complexity of audio extraction, video processing, and format conversion.

When to use?
  - When the subsystem is complex and the client shouldn't need to know its internal details.
  - When you want to define a clear entry point for a group of related operations.
  - When you want to isolate client code from changes in the subsystem.
  - When you're refactoring legacy code and want to wrap an old interface with a simpler one.

When NOT to use?
  - When the client needs access to detailed subsystem operations.
  - When the facade becomes a "god object" with dozens of methods.
  - When the subsystem is already simple and an extra layer doesn't simplify anything.

Tips and pitfalls:
  - Facade vs. Adapter: Adapter changes the interface of a single object, Facade creates
    a new simplified interface for an entire subsystem.
  - In the current implementation, NewMediaConverter() creates components internally;
    in more testable code you could accept interfaces as parameters.
  - Facade doesn't "hide" the subsystem -- it just offers a simpler API.
*/
package facade

import "fmt"

// AudioExtractor is responsible for extracting the audio track from a multimedia file.
type AudioExtractor struct{}

// ExtractAudio extracts the audio track from the given input.
func (a *AudioExtractor) ExtractAudio(input string) string {
	return fmt.Sprintf("audio_track(%s)", input)
}

// VideoProcessor is responsible for processing the video stream.
type VideoProcessor struct{}

// ProcessVideo processes the video from the given input.
func (v *VideoProcessor) ProcessVideo(input string) string {
	return fmt.Sprintf("processed_video(%s)", input)
}

// FormatConverter is responsible for converting and combining streams into the target format.
type FormatConverter struct{}

// Convert combines the video and audio streams into the given output format.
func (f *FormatConverter) Convert(video, audio, format string) string {
	return fmt.Sprintf("%s+%s.%s", video, audio, format)
}

// MediaConverter is a facade that coordinates audio extraction, video processing, and format conversion.
type MediaConverter struct {
	audio  *AudioExtractor
	video  *VideoProcessor
	format *FormatConverter
}

// NewMediaConverter creates a new facade with ready-made subsystem components.
func NewMediaConverter() *MediaConverter {
	return &MediaConverter{
		audio:  &AudioExtractor{},
		video:  &VideoProcessor{},
		format: &FormatConverter{},
	}
}

// Convert converts a multimedia file to the given format, coordinating the entire process under the hood.
func (m *MediaConverter) Convert(input, outputFormat string) string {
	audio := m.audio.ExtractAudio(input)
	video := m.video.ProcessVideo(input)
	result := m.format.Convert(video, audio, outputFormat)
	return result
}
