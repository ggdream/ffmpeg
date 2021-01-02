package ffmpeg

import (
	"bytes"
	"strings"
)



// M alias
type M			= map[string]interface{}

// Meta src info
type Meta struct {
	URI			string

	VCodec		string
	ACodec		string
}

// Options args options
type Options struct {
	Inputs		M
	Outputs		M	// string or Meta

	Quiet		bool
	Async		bool
}

// Handle join in the args
func (o *Options) Handle() string {
	var buffer	bytes.Buffer

	if o.Quiet {
		buffer.WriteString("-loglevel quiet ")
	}

	for k := range o.Inputs {
		buffer.WriteString("-i ")
		buffer.WriteString(k)
		buffer.WriteString(" ")
	}
	for k, v := range o.Outputs {
		buffer.WriteString(strings.Trim(v.(string), " "))
		buffer.WriteString(" ")
		buffer.WriteString(k)
		buffer.WriteString(" ")
	}

	return strings.TrimRight(buffer.String(), " ")
}
