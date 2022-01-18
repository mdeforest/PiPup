package accelerometer

import (
	"fmt"
	"math"
	"time"

	"github.com/mdeforest/PiPup/server/pkg/games"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

const kFilteringFactor float64 = 0.1
const sensitivityThreshold float64 = 1.0
const HasMoved = "has moved"

type Accelerometer struct {
	game    games.Game
	adaptor *raspi.Adaptor
	Driver  *i2c.MPU6050Driver
	Data    chan float64
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
			beforeAccelerometer := d.Accelerometer

			d.GetData()

			moved, vectorLength := hasMoved(beforeAccelerometer, d.Accelerometer)

			fmt.Printf("Moved: %t, Vector Length: %f\n", moved, vectorLength)

			//if moved {
			//	accelerometer.Data <- vectorLength
			//}
		})
	}

	robot := gobot.NewRobot(game.String(),
		[]gobot.Connection{a},
		[]gobot.Device{d},
		work)

	accelerometer.gobot = robot
	accelerometer.Data = make(chan float64)

	return accelerometer
}

func (a *Accelerometer) Start() error {
	if err := a.gobot.Start(); err != nil {
		return err
	}

	if err := a.Driver.Start(); err != nil {
		return err
	}

	return nil
}

func (a *Accelerometer) Stop() {
	a.gobot.Stop()
	a.Driver.Halt()
}

func hasMoved(beforeAccel i2c.ThreeDData, afterAccel i2c.ThreeDData) (bool, float64) {

	accelX := float64(afterAccel.X) - ((float64(afterAccel.X) * kFilteringFactor) + float64(beforeAccel.X)*(1.0-kFilteringFactor))
	accelY := float64(afterAccel.Y) - ((float64(afterAccel.Y) * kFilteringFactor) + float64(beforeAccel.Y)*(1.0-kFilteringFactor))
	accelZ := float64(afterAccel.Z) - ((float64(afterAccel.Z) * kFilteringFactor) + float64(beforeAccel.Z)*(1.0-kFilteringFactor))

	deltaX := math.Abs(accelX - float64(beforeAccel.X))
	deltaY := math.Abs(accelY - float64(beforeAccel.Y))
	deltaZ := math.Abs(accelZ - float64(beforeAccel.Z))

	bumpVectorLength := math.Sqrt(deltaX*deltaX + deltaY*deltaY + deltaZ*deltaZ)

	return bumpVectorLength > sensitivityThreshold, bumpVectorLength
}
