package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldVxByte struct {
	opcode chip8.Opcode
}

// LDVxByte - 6xkk - Set Vx = kk
func LDVxByte(opcode chip8.Opcode) Instruction {
	return &ldVxByte{opcode: opcode}
}

func (i *ldVxByte) Execute(c *chip8.Chip8) error {
	c.VX[i.opcode.Vx] = uint8(i.opcode.KK)
	c.ProgramCounter += 2

	return nil
}

func (i *ldVxByte) String() string {
	return fmt.Sprintf("LD Vx=0x%x byte=0x%x", i.opcode.Vx, i.opcode.KK)
}
