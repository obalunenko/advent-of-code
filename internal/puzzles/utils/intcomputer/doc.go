// Package intcomputer provides functionality for Intcode programs calculation
// Intcode programs are given as a list of integers; these values are used as the initial state for the computer's
// memory. When you run an Intcode program, make sure to start by initializing memory to the program's values.
// A position in memory is called an address (for example, the first value in memory is at "address 0").
//
// Opcodes (like 1, 2, or 99) mark the beginning of an instruction. The values used immediately after an opcode,
// f any, are called the instruction's parameters. For example, in the instruction 1,2,3,4, 1 is the opcode; 2, 3,
// and 4 are the parameters. The instruction 99 contains only an opcode and has no parameters.
//
// The address of the current instruction is called the instruction pointer; it starts at 0. After an instruction
// finishes, the instruction pointer increases by the number of values in the instruction; until you add more
// instructions to the computer, this is always 4 (1 opcode + 3 parameters) for the add and multiply instructions.
// (The halt instruction would increase the instruction pointer by 1, but it halts the program instead.)
//
// To run one, start by looking at the first integer (called position 0). Here, you will find an opcode -
// either 1, 2, or 99. The opcode indicates what to do; for example, 99 means that the program is finished and
// should immediately halt. Encountering an unknown opcode means something went wrong.
//
// Opcode 1 adds together numbers read from two positions and stores the result in a third position.
// The three integers immediately after the opcode tell you these three positions - the first two indicate
// the positions from which you should read the input values, and the third indicates the position at which the output
// should be stored.
//
// For example, if your Intcode computer encounters 1,10,20,30, it should read the values at positions 10 and 20,
// add those values, and then overwrite the value at position 30 with their sum.
//
// Opcode 2 works exactly like opcode 1, except it multiplies the two inputs instead of adding them. Again,
// the three integers after the opcode indicate where the inputs and outputs are, not their values.
//
// Once you're done processing an opcode, move to the next one by stepping forward 4 positions.
//
// For example, suppose you have the following program:
//
// 1,9,10,3,2,3,11,0,99,30,40,50
// For the purposes of illustration, here is the same program split into multiple lines:
//
// 1,9,10,3,
// 2,3,11,0,
// 99,
// 30,40,50
//
// The first four integers, 1,9,10,3, are at positions 0, 1, 2, and 3. Together, they represent the
// first opcode (1, addition), the positions of the two inputs (9 and 10), and the position of the output (3).
// To handle this opcode, you first need to get the values at the input positions: position 9 contains 30, and
// position 10 contains 40. Add these numbers together to get 70. Then, store this value at the output position;
// here, the output position (3) is at position 3, so it overwrites itself. Afterward, the program looks like this:
//
// 1,9,10,70,
// 2,3,11,0,
// 99,
// 30,40,50
// Step forward 4 positions to reach the next opcode, 2. This opcode works just like the previous, but it
// multiplies instead of adding. The inputs are at positions 3 and 11; these positions contain 70 and 50
// respectively. Multiplying these produces 3500; this is stored at position 0:
//
// 3500,9,10,70,
// 2,3,11,0,
// 99,
// 30,40,50
// Stepping forward 4 more positions arrives at opcode 99, halting the program.
package intcomputer
