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

type GLTextureData struct {
	ID     int64
	Width  int32
	Height int32
}

// Camera allows controlling the robot camera.
type Camera struct {
	c *camera.Module

	t token.Token
}

// StartVideo starts the camera video stream.
func (c *Camera) StartVideo() (err error) {
	endTrace := c.c.Logger().Trace("StartVideo")
	defer func() {
		endTrace("error", err)
	}()

	t, err := c.c.AddVideoCallback(func(frame *camera.RGB) {
		c.c.Logger().Error("AddVideoHandler: Unexpectedly got frame!")
	})

	if err != nil {
		c.t = t
	}

	return err
}

// StopVideo stops the camera video stream.
func (c *Camera) StopVideo() (err error) {
	endTrace := c.c.Logger().Trace("StopVideo")
	defer func() {
		if err == nil {
			c.t = 0
		}

		endTrace("error", err)
	}()

	return c.c.RemoveVideoCallback(c.t)
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

// RenderNextFrame requests the next frame to be rendered. This is used by iOS
// and the frame will be rendered to a texture associated with an OpenGLES 2.0
// context that was current when Start() is called. This should be called for
// for each frame to be rendered (up to 60 times per second).
func (c *Camera) RenderNextFrame() {
	c.c.RenderNextFrame()
}

// GLTextureData returns information about the current texture used for
// rendering frames. See RenderNextFrame() above.
func (c *Camera) GLTextureData() (*GLTextureData, error) {
	glTextureData, err := c.c.GLTextureData()
	if err != nil {
		return nil, err
	}

	return &GLTextureData{
		ID:     int64(glTextureData.ID),
		Width:  int32(glTextureData.Width),
		Height: int32(glTextureData.Height),
	}, nil
}
