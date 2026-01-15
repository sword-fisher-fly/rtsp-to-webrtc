package main

import (
	"os"

	"github.com/sword-fisher-fly/rtsp-to-webrtc/webrtc"
)

func main() {
	s, ok := webrtc.NewCore(os.Args[1:])
	if !ok {
		os.Exit(1)
	}
	s.Wait()
}
