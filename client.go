package mobile

import (
	"log/slog"

	"github.com/brunoga/robomaster"
	"github.com/brunoga/robomaster/module"
	"github.com/brunoga/robomaster/module/robot"
	"github.com/brunoga/robomaster/support/logger"
)

const (
	mobileModules = module.TypeAll
)

// Client is the main entry point for the mobile SDK.
type Client struct {
	l *logger.Logger
	c *robomaster.Client
}

// NewClient creates a new Client instance. If appID is 0, the client will try
// to connect to the first available Robomaster robot. If it is non-zero, it
// will only connect to a robot that is broadcasting the given appID. The appID
// can be configured in the robot through a qrcode.
func NewClient(appID int64) (*Client, error) {
	l := logger.New(slog.LevelDebug)

	c, err := robomaster.NewWithModules(l, uint64(appID), mobileModules)
	if err != nil {
		return nil, err
	}

	return &Client{
		l: l,
		c: c,
	}, nil
}

// NewWifiDirectClient creates a new Client instance that will connect to a
// Robomaster robot using Wifi Direct.
func NewWifiDirectClient() (*Client, error) {
	l := logger.New(logger.LevelTrace)

	c, err := robomaster.NewWifiDirectWithModules(l, mobileModules)
	if err != nil {
		return nil, err
	}

	return &Client{
		c: c,
		l: l,
	}, nil
}

// Start starts the client.
func (c *Client) Start() error {
	// Start underlying client.
	err := c.c.Start()
	if err != nil {
		return err
	}

	// Enable movement control (i.e. controller.Move()).
	return c.c.Robot().EnableFunction(robot.FunctionTypeMovementControl, true)
}

// Camera returns the Camera instance for the client.
func (c *Client) Camera() *Camera {
	return &Camera{
		c: c.c.Camera(),
	}
}

// Controller returns the Controller instance for the client.
func (c *Client) Chassis() *Chassis {
	return &Chassis{
		// Keeping compatibilitty after refactoring.
		c: c.c.Controller(),
	}
}

// Connnection returns the Connection instance for the client.
func (c *Client) Connection() *Connection {
	return &Connection{
		c: c.c.Connection(),
	}
}

// GamePad returns the GamePad instance for the client. The GamePad is optional
// and may be nil.
func (c *Client) GamePad() *GamePad {
	return &GamePad{
		g: c.c.GamePad(),
	}
}

// Robot returns the Robot instance for the client.
func (c *Client) Robot() *Robot {
	return &Robot{
		r: c.c.Robot(),
	}
}

// Stop stops the client.
func (c *Client) Stop() error {
	return c.c.Stop()
}
