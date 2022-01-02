package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type jpAddr struct {
	opcode chip8.Opcode
}

// JPAddr - 1nnn - Jump to location nnn
func JPAddr(opcode chip8.Opcode) Instruction {
	return &jpAddr{opcode: opcode}
}

func (i *jpAddr) Execute(c *chip8.Chip8) error {
	c.ProgramCounter = i.opcode.NNN

	return nil
}

func (i *jpAddr) String() string {
	return fmt.Sprintf("JP addr=0x%x", i.opcode.NNN)
}
