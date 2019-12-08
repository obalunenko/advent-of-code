package fuel

import (
	"bufio"
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type module struct {
	mass string
}

func (m module) fuel() (int, error) {
	mass, err := strconv.Atoi(m.mass)
	if err != nil {
		return 0, err
	}

	diff := mass % 3
	if diff != 0 {
		mass = mass - diff
	}

	f := (mass / 3) - 2

	return f, nil
}

func calc(in chan input, res chan result) {
	for i := range in {
		go fuelForModule(i.m, res)
	}
}

func fuelForModule(m module, res chan result) {
	f, err := m.fuel()
	res <- result{
		fuel: f,
		err:  err,
	}
}

type input struct {
	m module
}

type result struct {
	fuel int
	err  error
}

func calculate(filepath string) (int, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return 0, errors.Wrap(err, "failed to open file")
	}

	defer file.Close()

	var sum int

	in := make(chan input)
	res := make(chan result)

	go calc(in, res)

	var lines int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		in <- input{
			module{
				mass: scanner.Text(),
			},
		}
		lines++
	}

	close(in)

	if err = scanner.Err(); err != nil {
		return 0, errors.Wrap(err, "scanner error")
	}

	for lines > 0 {
		r := <-res
		if r.err != nil {
			return 0, err
		}

		sum += r.fuel

		lines--
	}

	close(res)

	return sum, nil
}
