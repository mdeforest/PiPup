package accelerometer

import (
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

	accelerometer.Driver.AddEvent(HasMoved)

	work := func() {
		gobot.Every(100*time.Millisecond, func() {
			beforeAccelerometer := d.Accelerometer

			d.GetData()

			moved, vectorLength := hasMoved(beforeAccelerometer, d.Accelerometer)

			if moved {
				accelerometer.Driver.Publish(HasMoved, vectorLength)
			}
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
