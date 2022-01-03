package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type shrVxVyLeastSignificant struct {
	opcode chip8.Opcode
}

// SHRVxVyLeastSignificant - 8xy6 - Set Vx = Vx SHR 1 if the least-significant bit of Vx is 1.
func SHRVxVyLeastSignificant(opcode chip8.Opcode) Instruction {
	return &shrVxVyLeastSignificant{opcode: opcode}
}

func (i *shrVxVyLeastSignificant) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]

	c.VX[0xF] = xValue & 0x1
	c.VX[i.opcode.Vx] = xValue >> 1

	c.ProgramCounter += 2

	return nil
}

func (i *shrVxVyLeastSignificant) String() string {
	return fmt.Sprintf("SHR Vx=0x%x", i.opcode.Vx)
}
