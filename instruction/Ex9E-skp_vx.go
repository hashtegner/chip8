package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type skpVx struct {
	opcode chip8.Opcode
}

// SkpVx - Ex9E - Ex9E - Skip next instruction if key with the value of Vx is pressed.
func SkpVx(opcode chip8.Opcode) Instruction {
	return &skpVx{opcode: opcode}
}

func (i *skpVx) Execute(c *chip8.Chip8) error {
	value := c.VX[i.opcode.Vx]
	key := c.Keys[value]
	if key > 0 {
		c.ProgramCounter += 4
	} else {
		c.ProgramCounter += 2
	}

	return nil
}

func (i *skpVx) String() string {
	return fmt.Sprintf("SKP Vx=0x%x", i.opcode.Vx)
}
