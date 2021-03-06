package levels

import (
	"fmt"
	"time"

	"github.com/mdeforest/PiPup/server/internal/app/pod/services/dispenser"
	"github.com/mdeforest/PiPup/server/pkg/accelerometer"
	"github.com/mdeforest/PiPup/server/pkg/score"
)

type levelOne struct {
}

func (l Level) NewLevelOne() *levelOne {
	return &levelOne{}
}

func (l *levelOne) PlayLevel(accelerometer *accelerometer.Accelerometer, s *score.Score, length int) {
	// just wait for accelerometer feedback
	select {
	case <-accelerometer.Data:
		fmt.Println("moved")
		s.IncreaseScore(1)
		dispenser.DispenseTreats()
		return
	case <-time.After(time.Duration(length) * time.Second):
		fmt.Println("timeout")
		return
	}
}
