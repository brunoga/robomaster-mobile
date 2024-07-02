package mobile

import "github.com/brunoga/robomaster/module/sdcard"

// SDCard allows handling an SD card inserted into the robot.
type SDCard struct {
	m *sdcard.Module
}

// IsInserted returns true if an SD card is inserted into the robot.
func (s *SDCard) IsInserted() (bool, error) {
	return s.m.IsInserted()
}

// Format formats the SD card.
func (s *SDCard) Format() error {
	return s.m.Format()
}

// IsFormatting returns true if the SD card is being formatted.
func (s *SDCard) IsFormatting() (bool, error) {
	return s.m.IsFormatting()
}

// IsFull returns true if the SD card is full.
func (s *SDCard) IsFull() (bool, error) {
	return s.m.IsFull()
}

// HasError returns true if the SD card has an error.
func (s *SDCard) HasError() (bool, error) {
	return s.m.HasError()
}

// TotalSpaceInMB returns the total space in MB of the SD card.
func (s *SDCard) TotalSpaceInMB() (int64, error) {
	space, err := s.m.TotalSpaceInMB()

	return int64(space), err
}

// RemainingSpaceInMB returns the remaining space in MB of the SD card.
func (s *SDCard) RemainingSpaceInMB() (int64, error) {
	space, err := s.m.RemainingSpaceInMB()

	return int64(space), err
}

// AvailablePhotoCount returns the estimated available photo count in the SD
// card.
func (s *SDCard) AvailablePhotoCount() (int64, error) {
	count, err := s.m.AvailablePhotoCount()

	return int64(count), err
}

// AvailableRecordingTimeInSeconds returns the estimated available recording
// time in seconds.
func (s *SDCard) AvailableRecordingTimeInSeconds() (int64, error) {
	time, err := s.m.AvailableRecordingTimeInSeconds()

	return int64(time), err
}
