package fuel

import (
	"bufio"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type module struct {
	mass int
}

func (m module) fuel() int {
	mass := m.mass

	diff := mass % 3
	if diff != 0 {
		mass = mass - diff
	}

	f := (mass / 3) - 2

	return f
}

func calc(in chan module, res chan int, done chan struct{}) {
	for i := range in {
		go fuelForModule(i, res, done)
	}
}

func fuelForModule(m module, res chan int, done chan struct{}) {
	var isDone bool

	for !isDone {
		f := m.fuel()

		res <- f

		if f/3 > 1 {
			m.mass = f
		} else {
			isDone = true
		}
	}

	done <- struct{}{}
}

func calculate(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, errors.Wrap(err, "failed to open file")
	}

	defer file.Close()

	var sum int

	in := make(chan module)
	res := make(chan int)
	done := make(chan struct{})

	go calc(in, res, done)

	var (
		lines int
		mass  int
	)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mass, err = strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, err
		}

		in <- module{
			mass: mass,
		}
		lines++
	}

	close(in)

	if err = scanner.Err(); err != nil {
		return 0, errors.Wrap(err, "scanner error")
	}

	for lines > 0 {
		select {
		case r := <-res:
			sum += r
		case <-done:
			lines--
		}
	}

	close(res)

	return sum, nil
}
