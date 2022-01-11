package levels

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/divan/num2words"
	"github.com/mdeforest/PiPup/server/pkg/accelerometer"
	"github.com/mdeforest/PiPup/server/pkg/score"
)

type Playable interface {
	PlayLevel(accelerometer *accelerometer.Accelerometer, s *score.Score, length int) int
}

type Level struct {
	num int
}

func CreateLevel(num int) (Playable, error) {
	_, filename, _, ok := runtime.Caller(1)

	if !ok {
		return nil, errors.New("could not get number of levels")
	}

	files, err := ioutil.ReadDir(filepath.Dir(filename))

	if err != nil {
		return nil, err
	}

	for i := 1; i < len(files); i++ {
		if num == i {
			words := strings.Title(num2words.Convert(num))

			level := Level{num: num}

			funcName := "NewLevel" + words

			method := reflect.ValueOf(level).MethodByName(funcName)

			if !method.IsValid() {
				return nil, errors.New("not a valid level")
			}

			result := method.Call(nil)
			ret := result[0].Interface().(Playable)

			return ret, nil
		}
	}

	return nil, errors.New("not a valid level")
}
