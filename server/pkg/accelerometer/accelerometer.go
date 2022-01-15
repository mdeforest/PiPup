package accelerometer

import (
	"fmt"
	"time"

	"github.com/mdeforest/PiPup/server/pkg/games"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

type Accelerometer struct {
	game    games.Game
	adaptor *raspi.Adaptor
	Driver  *i2c.MPU6050Driver
	gobot   *gobot.Robot
}

func NewAccelerometer(game games.Game) *Accelerometer {
	a := raspi.NewAdaptor()
	d := i2c.NewMPU6050Driver(a)

	accelerometer := &Accelerometer{
		game:    game,
		adaptor: a,
		Driver:  d,
	}

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			d.GetData()

			fmt.Println("Accelerometer", d.Accelerometer)
			fmt.Println("Gyroscope", d.Gyroscope)
			fmt.Println("Temperature", d.Temperature)
		})
	}

	robot := gobot.NewRobot(game.String(),
		[]gobot.Connection{a},
		[]gobot.Device{d},
		work)

	accelerometer.gobot = robot

	return accelerometer
}

func (a *Accelerometer) Start() error {
	return a.gobot.Start()
}

func (a *Accelerometer) Stop() {
	a.gobot.Stop()
}
