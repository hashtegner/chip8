package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type addIVx struct {
	opcode chip8.Opcode
}

// AddIVx - Fx1E - Set I = I + Vx
func AddIVx(opcode chip8.Opcode) Instruction {
	return &addIVx{opcode: opcode}
}

func (i *addIVx) Execute(c *chip8.Chip8) error {
	xValue := c.VX[i.opcode.Vx]
	c.Index += uint16(xValue)

	c.ProgramCounter += 2

	return nil
}

func (i *addIVx) String() string {
	return fmt.Sprintf("SKNP Vx=0x%x", i.opcode.Vx)
}
