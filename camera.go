package mobile

import (
	"github.com/brunoga/robomaster/support/token"

	"github.com/brunoga/robomaster/module/camera"
)

const (
	CameraHorizontalResolutionPoints = camera.HorizontalResolutionPoints
	CameraVerticalResolutionPoints   = camera.VerticalResolutionPoints
	CameraHorizontalFOVDegrees       = camera.HorizontalFOVDegrees
	CameraVerticalFOVDegrees         = camera.VerticalFOVDegrees
)

// VideoHandler is the interface that must be implemented by types that want
// to handle video frames from the camera.
type VideoHandler interface {
	// HandleVideo is called when a new video frame is received from the camera.
	// rgb24Data, as the name name implies, is the raw data for a RGB24 image and
	// its dimensions are 1280x720.
	HandleVideo(rgb24Data []byte)
}

// Camera allows controlling the robot camera.
type Camera struct {
	c *camera.Camera
}

// AddVideoHandler adds a new video handler to the camera. If this is the first
// video handler added, the camera will start sending video frames.
func (c *Camera) AddVideoHandler(handler VideoHandler) (token int64, err error) {
	endTrace := c.c.Logger().Trace("AddVideoHandler", "handler", handler)
	defer func() {
		endTrace("token", token, "error", err)
	}()

	c.c.Logger().Debug("AddVideoHandler", "handler.handleVideo", handler.HandleVideo)

	t, err := c.c.AddVideoCallback(func(frame *camera.RGB) {
		c.c.Logger().Debug("AddVideoHandler: Got frame!")
		handler.HandleVideo(frame.Pix)
	})

	return int64(t), err
}

// RemoveVideoHandler removes a video handler from the camera. If this is the
// last video handler removed, the camera will stop sending video frames.
func (c *Camera) RemoveVideoHandler(t int64) (err error) {
	endTrace := c.c.Logger().Trace("RemoveVideoHandler", "token", t)
	defer func() {
		endTrace("error", err)
	}()

	return c.c.RemoveVideoCallback(token.Token(t))
}

// StartRecordingVideo starts recording video from the camera to the robot's
// SD card.
func (c *Camera) StartRecordingVideo() (err error) {
	endTrace := c.c.Logger().Trace("StartRecordingVideo")
	defer func() {
		endTrace("error", err)
	}()

	return c.c.StartRecordingVideo()
}

// IsRecordingVideo returns true if the camera is currently recording video.
func (c *Camera) IsRecordingVideo() (isRecording bool, err error) {
	endTrace := c.c.Logger().Trace("IsRecordingVideo")
	defer func() {
		endTrace("isRecording", isRecording, "error", err)
	}()

	return c.c.IsRecordingVideo()
}

// RecordingTimeInSeconds returns the current recording time in seconds.
func (c *Camera) RecordingTimeInSeconds() (recordingTime int64) {
	endTrace := c.c.Logger().Trace("RecordingTimeInSeconds")
	defer func() {
		endTrace("recordingTime", recordingTime)
	}()

	return int64(c.c.RecordingTime().Seconds())
}

// StopRecordingVideo stops recording video from the camera to the robot's SD
// card.
func (c *Camera) StopRecordingVideo() (err error) {
	endTrace := c.c.Logger().Trace("StopRecordingVideo")
	defer func() {
		endTrace("error", err)
	}()

	return c.c.StopRecordingVideo()
}
