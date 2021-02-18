package day01

type move string

const (
	up   move = "("
	down move = ")"
)

type floor int

const (
	ground   floor = 0
	basement floor = -1
)

type elevator struct {
	floor floor
}

func newElevator() *elevator {
	return &elevator{
		floor: ground,
	}
}

func (e *elevator) Up() {
	e.floor++
}

func (e *elevator) Down() {
	e.floor--
}

func (e elevator) Floor() floor {
	return e.floor
}

func (e *elevator) Move(m move) {
	switch m {
	case up:
		e.Up()
	case down:
		e.Down()
	}
}
