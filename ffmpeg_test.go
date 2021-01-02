package ffmpeg

import (
	"fmt"
	"testing"
)


func TestOptions(t *testing.T) {
	o := Options{
		Inputs: M{
			"video.mp4": nil,
			"audio.mp3": nil,
		},
		Outputs: M{
			"result.flv": "-c:v copy -c:a copy",
		},
		Quiet: true,
	}

	fmt.Println(o.Handle())
}


// func TestNew(t *testing.T) {
// 	New(&Options{
// 		Inputs: M{
// 			"./src/1.mp4": nil,
// 			"./src/1.mp3": nil,
// 		},
// 		Outputs: M{
// 			"res.mp4": "-c:v libx264 -c:a copy",
// 		},
// 	}).Run()

// 	fmt.Println("sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss")
// }
