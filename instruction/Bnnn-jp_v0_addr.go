package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type jpV0Addr struct {
	opcode chip8.Opcode
}

// JpV0Addr - Bnnn - Jump to location nnn + v0
func JpV0Addr(opcode chip8.Opcode) Instruction {
	return &jpV0Addr{opcode: opcode}
}

func (i *jpV0Addr) Execute(c *chip8.Chip8) error {
	c.ProgramCounter = i.opcode.NNN + uint16(c.VX[0])

	return nil
}

func (i *jpV0Addr) String() string {
	return fmt.Sprintf("JP V0 addr=0x%x", i.opcode.NNN)
}
