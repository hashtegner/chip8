package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type callAddr struct {
	opcode chip8.Opcode
}

// CallAddr - 2nnn - Call subroutine at nnn
func CallAddr(opcode chip8.Opcode) Instruction {
	return &callAddr{opcode: opcode}
}

func (i *callAddr) Execute(c *chip8.Chip8) error {
	c.Stack[c.StackPointer] = c.ProgramCounter + 2
	c.StackPointer += 1
	c.ProgramCounter = i.opcode.NNN

	return nil
}

func (i *callAddr) String() string {
	return fmt.Sprintf("CALL addr=0x%x", i.opcode.NNN)
}
