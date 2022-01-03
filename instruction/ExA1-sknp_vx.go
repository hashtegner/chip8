package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type sknpVx struct {
	opcode chip8.Opcode
}

// SknpVx - Ex9E - Skip next instruction if key with the value of Vx is not pressed.
func SknpVx(opcode chip8.Opcode) Instruction {
	return &sknpVx{opcode: opcode}
}

func (i *sknpVx) Execute(c *chip8.Chip8) error {
	value := c.VX[i.opcode.Vx]
	key := c.Keys[value]
	if key == 0 {
		c.ProgramCounter += 4
	} else {
		c.ProgramCounter += 2
	}

	return nil
}

func (i *sknpVx) String() string {
	return fmt.Sprintf("SKNP Vx=0x%x", i.opcode.Vx)
}
