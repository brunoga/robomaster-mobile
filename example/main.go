package main

import (
	"fmt"
	"time"

	mobile "github.com/brunoga/robomaster-mobile"
)

type VideoHandler struct{}

func (vh *VideoHandler) HandleVideo(rgb24Data []byte) {
	// Do something with the video frame.
	fmt.Println("Received video frame.")
}

func main() {
	c, err := mobile.NewClient(0)
	if err != nil {
		panic(err)
	}

	err = c.Start()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = c.Stop()
		if err != nil {
			panic(err)
		}
	}()

	camera := c.Camera()

	videoHandler := &VideoHandler{}

	// Add a video handler.
	t, err := camera.AddVideoHandler(videoHandler)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = camera.RemoveVideoHandler(t)
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(5 * time.Second)
}
