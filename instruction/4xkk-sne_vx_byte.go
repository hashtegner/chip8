package instruction

import (
	"fmt"

	"github.com/hashtegner/chip8/chip8"
)

type sneVxByte struct {
	opcode chip8.Opcode
}

// SNEVXByte - 4xkk - Skip next instruction if Vx != kk
func SNEVxByte(opcode chip8.Opcode) Instruction {
	return &sneVxByte{opcode: opcode}
}

func (i *sneVxByte) Execute(c *chip8.Chip8) error {
	value := c.VX[i.opcode.Vx]

	if value != uint8(i.opcode.KK) {
		c.ProgramCounter += 4
	} else {
		c.ProgramCounter += 2
	}

	return nil

}

func (i *sneVxByte) String() string {
	return fmt.Sprintf("SNE Vx=0x%x byte=0x%x", i.opcode.Vx, i.opcode.KK)
}
