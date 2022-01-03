package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type ldVxK struct {
	opcode chip8.Opcode
}

// LDVxK - Fx0A - Wait for key press, store the value of the key in Vx
func LDVxK(opcode chip8.Opcode) Instruction {
	return &ldVxK{opcode: opcode}
}

func (i *ldVxK) Execute(c *chip8.Chip8) error {
loop:
	for {
		for j := 0; j < len(c.Keys); j++ {
			value := c.Keys[j]
			if value > 0 {
				c.VX[i.opcode.Vx] = uint8(j)
				break loop
			}
		}
	}

	c.ProgramCounter += 2

	return nil
}

func (i *ldVxK) String() string {
	return fmt.Sprintf("LD Vx=0x%x K", i.opcode.Vx)
}
