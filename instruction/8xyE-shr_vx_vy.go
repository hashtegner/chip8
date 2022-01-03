package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type shrVxVyMostSignificant struct {
	opcode chip8.Opcode
}

// SHRVxVyMostSignificant - 8xyE - Set Vx = Vx SHR 1 if the most-significant bit of Vx is 1.
func SHRVxVyMostSignificant(opcode chip8.Opcode) Instruction {
	return &shrVxVyMostSignificant{opcode: opcode}
}

func (i *shrVxVyMostSignificant) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]

	c.VX[0xF] = (xValue >> 7) & 0x1
	c.VX[i.opcode.Vx] = xValue << 1

	c.ProgramCounter += 2

	return nil
}

func (i *shrVxVyMostSignificant) String() string {
	return fmt.Sprintf("SHL Vx=0x%x Vy=0x%x", i.opcode.Vx, i.opcode.Vy)
}
