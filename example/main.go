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

	// Add a video handler.
	err = camera.StartVideo()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = camera.StopVideo()
		if err != nil {
			panic(err)
		}
	}()

	time.Sleep(5 * time.Second)
}
