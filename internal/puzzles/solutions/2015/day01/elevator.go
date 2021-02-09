package day01

type move string

const (
	up   move = "("
	down move = ")"

	basement = -1
)

type elevator struct {
	floor int
}

func newElevator() *elevator {
	return &elevator{
		floor: 0,
	}
}

func (e *elevator) Up() {
	e.floor++
}

func (e *elevator) Down() {
	e.floor--
}

func (e elevator) Floor() int {
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
