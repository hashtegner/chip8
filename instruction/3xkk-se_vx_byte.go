package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type seVxByte struct {
	opcode chip8.Opcode
}

// SEVXByte - 3xkk - Skip next instruction if Vx == kk
func SEVxByte(opcode chip8.Opcode) Instruction {
	return &seVxByte{opcode: opcode}
}

func (i *seVxByte) Execute(c *chip8.Chip8) error {
	value := c.VX[i.opcode.Vx]

	if value == uint8(i.opcode.KK) {
		c.ProgramCounter += 4
	} else {
		c.ProgramCounter += 2
	}

	return nil
}

func (i *seVxByte) String() string {
	return fmt.Sprintf("SE Vx=0x%x byte=0x%x", i.opcode.Vx, i.opcode.KK)
}
