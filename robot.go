package mobile

import "github.com/brunoga/robomaster/module/robot"

// Robot allows reading the robot parameters.
type Robot struct {
	r *robot.Robot
}

// BatteryPowerPercent returns the current battery power percent (0 to 100).
func (r *Robot) BatteryPowerPercent() int8 {
	return int8(r.r.BatteryPowerPercent())
}

// ChassisSpeedLevel returns the current chassis speed level.
func (r *Robot) ChassisSpeedLevel() (int8, error) {
	chassisSpeedLevel, err := r.r.ChassisSpeedLevel()
	if err != nil {
		return -1, err
	}

	return int8(chassisSpeedLevel), nil
}

// SetChassisSpeedLevel sets the chassis speed level. Valid values are 0 (Fast),
// 1 (Medium) and 2 (Slow).
func (r *Robot) SetChassisSpeedLevel(speedLevel int8) error {
	return r.r.SetChassisSpeedLevel(robot.ChassisSpeedLevel(speedLevel))
}
