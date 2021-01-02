package ffmpeg

import (
	"io"
	"os/exec"
	"strings"
)


// FFmpeg ...
type FFmpeg struct {
	async		bool
	writer		io.Writer

	Raw			string
}

// SetWriter ...
func (f *FFmpeg) SetWriter(writer io.Writer){
	f.writer = writer
}

// Run ...
func (f *FFmpeg) Run() error {
	cmd := exec.Command("ffmpeg", strings.Split(f.Raw, " ")...)
	cmd.Stderr = f.writer

	if f.async {
		return cmd.Start()	// Async
	}

	return cmd.Run()		// Sync
}
