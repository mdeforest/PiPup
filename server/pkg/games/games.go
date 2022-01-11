package games

type Game int64

const (
	Undefined Game = iota
	Pod
)

func (g Game) String() string {
	switch g {
	case Pod:
		return "Pod"
	}

	return "unknown"
}
