package pod

import (
	"fmt"
	"sync"

	"github.com/mdeforest/PiPup/server/pkg/accelerometer"
	"github.com/mdeforest/PiPup/server/pkg/games"
	"github.com/mdeforest/PiPup/server/pkg/games/pod/levels"
	"github.com/mdeforest/PiPup/server/pkg/score"
)

type PodGame struct {
	level      int
	length     int
	scoreboard *score.Score
}

var game *PodGame

// level (integer)
func NewPodGame(level int, length int) *PodGame {
	if game == nil {
		game = &PodGame{
			scoreboard: score.NewScore(level),
		}
	}

	game.SetLevel(level)
	game.SetLength(length)

	return game
}

func (g *PodGame) SetLevel(level int) {
	g.level = level
}

func (g *PodGame) SetLength(length int) {
	g.length = length
}

func (g *PodGame) Start(wg *sync.WaitGroup) {
	//defer wg.Done()

	fmt.Println("Game Started")

	gameLevel, err := levels.CreateLevel(g.level)

	if err != nil {
		return
	}

	fmt.Println("Create accelerator")

	accelerometer := accelerometer.NewAccelerometer(games.Pod)
	if err := accelerometer.Start(); err != nil {
		return
	}

	fmt.Println("accelerometer started")

	gameLevel.PlayLevel(accelerometer, g.scoreboard, g.length)

	accelerometer.Stop()

	fmt.Println("Game Finished")
}

func (g *PodGame) Reset() {
	g = nil
	game = nil
}
