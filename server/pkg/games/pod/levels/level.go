package levels

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/divan/num2words"
	"github.com/mdeforest/PiPup/server/pkg/accelerometer"
	"github.com/mdeforest/PiPup/server/pkg/score"
)

type Playable interface {
	PlayLevel(accelerometer *accelerometer.Accelerometer, s *score.Score, length int)
}

type Level struct {
	num int
}

func CreateLevel(num int) (Playable, error) {
	words := strings.Title(num2words.Convert(num))

	level := Level{num: num}

	funcName := "NewLevel" + words

	method := reflect.ValueOf(level).MethodByName(funcName)

	if !method.IsValid() {
		return nil, errors.New("not a valid level")
	}

	fmt.Println("new level")

	result := method.Call(nil)
	ret := result[0].Interface().(Playable)

	return ret, nil
}
