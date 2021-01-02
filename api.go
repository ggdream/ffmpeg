package ffmpeg

import (
	"os"

)


// New use ffmpeg by `options`
func New(options *Options) *FFmpeg {
	raw := options.Handle()
	return &FFmpeg{
		Raw: raw,
		writer: os.Stderr,
		async: options.Async,
	}
}

// Raw use ffmpeg by `cmd raw code`
func Raw(raw string) *FFmpeg {
	return &FFmpeg{
		Raw: raw,
		writer: os.Stderr,
	}
}
