package accelerometer

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/mdeforest/PiPup/server/pkg/games"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

const sensitivityThreshold float64 = 200.0
const HasMoved = "has moved"

type Accelerometer struct {
	game    games.Game
	adaptor *raspi.Adaptor
	Driver  *i2c.MPU6050Driver
	Data    chan float64
	gobot   *gobot.Robot
	work    *gobot.RobotWork
}

func NewAccelerometer(game games.Game) *Accelerometer {
	a := raspi.NewAdaptor()
	d := i2c.NewMPU6050Driver(a)

	accelerometer := &Accelerometer{
		game:    game,
		adaptor: a,
		Driver:  d,
	}

	robot := gobot.NewRobot(game.String(),
		[]gobot.Connection{a},
		[]gobot.Device{d})

	d.GetData()

	work := robot.Every(context.Background(), 100*time.Millisecond, func() {
		beforeAccelerometer := d.Accelerometer

		d.GetData()

		fmt.Println(d.Accelerometer)

		moved, vectorLength := hasMoved(beforeAccelerometer, d.Accelerometer)

		fmt.Printf("Moved: %t, Vector Length: %f\n", moved, vectorLength)

		if moved {
			accelerometer.Data <- vectorLength
		}
	})

	accelerometer.gobot = robot
	accelerometer.Data = make(chan float64)
	accelerometer.work = work

	return accelerometer
}

func (a *Accelerometer) Start() error {
	if err := a.gobot.Start(); err != nil {
		return err
	}

	return nil
}

func (a *Accelerometer) Stop() {
	a.work.CallCancelFunc()
	a.Driver.Halt()
	a.gobot.Stop()
}

func hasMoved(beforeAccel i2c.ThreeDData, afterAccel i2c.ThreeDData) (bool, float64) {
	deltaX := float64(afterAccel.X - beforeAccel.X)
	deltaY := float64(afterAccel.Y - beforeAccel.Y)
	deltaZ := float64(afterAccel.Z - beforeAccel.Z)

	fmt.Println(deltaX, deltaY, deltaZ)

	if math.Abs(deltaX) > sensitivityThreshold {
		return true, deltaX
	} else if math.Abs(deltaY) > sensitivityThreshold {
		return true, deltaY
	} else if math.Abs(deltaZ) > sensitivityThreshold {
		return true, deltaZ
	}

	return false, 0
}
