// Package intcomputer provides functionality for Intcode programs calculation
// Intcode programs are given as a list of integers; these values are used as the initial state for the computer's
// memory. When you run an Intcode program, make sure to start by initializing memory to the program's values.
// A position in memory is called an address (for example, the first value in memory is at "address 0").
//
// Opcodes (like 1, 2, or 99) mark the beginning of an instruction. The values used immediately after an opcode,
// if any, are called the instruction's parameters. For example, in the instruction 1,2,3,4, 1 is the opcode; 2, 3,
// and 4 are the parameters. The instruction 99 contains only an opcode and has no parameters.
package intcomputer

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// IntComputer represents inctomputer instance.
type IntComputer struct {
	memory  map[int]int
	initial []int
}

const (
	optAdd   = 1
	optMult  = 2
	optAbort = 99

	shift = 4
)

// New creates instance of IntComputer from passed intcode program.
func New(in io.Reader) (IntComputer, error) {
	var c IntComputer

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(in); err != nil {
		return c, fmt.Errorf("failed to read: %w", err)
	}

	raw := strings.TrimSpace(buf.String())

	nums := strings.Split(raw, ",")
	c.initial = make([]int, len(nums))
	c.memory = make(map[int]int, len(nums))

	for i, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			return c, fmt.Errorf("failed to convert string to int: %w", err)
		}

		c.initial[i] = n
		c.memory[i] = n
	}

	return c, nil
}

// Execute executes intcode program that was loaded to IntComputer.
func (c *IntComputer) Execute() (int, error) {
	var (
		result int
		err    error
	)

loop:
	for i := 0; i < len(c.memory); i += shift {
		opt, aPos, bPos, resPos := c.memory[i], c.memory[i+1], c.memory[i+2], c.memory[i+3]
		switch opt {
		case optAdd:
			if err = c.add(aPos, bPos, resPos); err != nil {
				return 0, fmt.Errorf("failed to add pos[%d]: [intcode:%d %d %d %d]: %w",
					i, opt, aPos, bPos, resPos, err)
			}
		case optMult:
			if err = c.mult(aPos, bPos, resPos); err != nil {
				return 0, fmt.Errorf("failed to mult pos[%d]: [intcode:%d %d %d %d]: %w",
					i, opt, aPos, bPos, resPos, err)
			}
		case optAbort:
			result, err = c.abort()

			break loop
		default:
			result = -1
			err = fmt.Errorf("not supported opt code [%d] at pos [%d]", opt, i)

			break loop
		}
	}

	return result, err
}

func (c *IntComputer) add(aPos, bPos, resPos int) error {
	a, ok := c.memory[aPos]
	if !ok {
		return fmt.Errorf("value not exist [apos:%d]", aPos)
	}

	b, ok := c.memory[bPos]
	if !ok {
		return fmt.Errorf("value not exist [bpos:%d]", bPos)
	}

	res := a + b
	c.memory[resPos] = res

	return nil
}

func (c *IntComputer) mult(aPos, bPos, resPos int) error {
	a, ok := c.memory[aPos]
	if !ok {
		return errors.New("value not exist")
	}

	b, ok := c.memory[bPos]
	if !ok {
		return errors.New("value not exist")
	}

	res := a * b
	c.memory[resPos] = res

	return nil
}

func (c *IntComputer) abort() (int, error) {
	res, ok := c.memory[0]
	if !ok {
		return 0, errors.New("value not exist")
	}

	return res, nil
}

// Input allow to input noun and verb into intcode program for execution.
// noun - 2nd position in intcode;
// verb - 3rd position in intcode.
func (c *IntComputer) Input(noun, verb int) {
	c.memory[1] = noun
	c.memory[2] = verb
}

// Reset resets computer's memory to the initial state.
func (c *IntComputer) Reset() {
	c.memory = make(map[int]int, len(c.initial))

	for i, n := range c.initial {
		n := n
		c.memory[i] = n
	}
}
