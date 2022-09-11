package crosswd

type WordDirection int

const (
	Horizontal WordDirection = 0
	Vertical   WordDirection = 1
)

func (d WordDirection) String() string {
	switch d {
	case Horizontal:
		return "Horizontal"
	case Vertical:
		return "Vertical"
	default:
		return "unknown"
	}
}

func (d WordDirection) Inverted() WordDirection {
	return WordDirection((int(d) + 1) % 2)
}
