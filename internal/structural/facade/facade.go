package facade

import "fmt"

type AudioExtractor struct{}

func (a *AudioExtractor) ExtractAudio(input string) string {
	return fmt.Sprintf("audio_track(%s)", input)
}

type VideoProcessor struct{}

func (v *VideoProcessor) ProcessVideo(input string) string {
	return fmt.Sprintf("processed_video(%s)", input)
}

type FormatConverter struct{}

func (f *FormatConverter) Convert(video, audio, format string) string {
	return fmt.Sprintf("%s+%s.%s", video, audio, format)
}

type MediaConverter struct {
	audio  *AudioExtractor
	video  *VideoProcessor
	format *FormatConverter
}

func NewMediaConverter() *MediaConverter {
	return &MediaConverter{
		audio:  &AudioExtractor{},
		video:  &VideoProcessor{},
		format: &FormatConverter{},
	}
}

func (m *MediaConverter) Convert(input, outputFormat string) string {
	audio := m.audio.ExtractAudio(input)
	video := m.video.ProcessVideo(input)
	result := m.format.Convert(video, audio, outputFormat)
	return result
}
