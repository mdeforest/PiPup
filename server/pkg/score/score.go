package score

import (
	"fmt"
	"strconv"
)

type Score struct {
	points int
	hasWon bool
	level  int
}

func NewScore(level int) *Score {
	return &Score{
		points: 0,
		hasWon: false,
		level:  level,
	}
}

func (s *Score) IncreaseScore(points int) {
	fmt.Println(strconv.Itoa(points) + " points scored")
	s.points += points
}

func (s *Score) WinGame() {
	s.hasWon = true
}
