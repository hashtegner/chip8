package instruction

import (
	"fmt"
	"math/rand"

	"github.com/hashtegner/chip8/chip8"
)

type rndVxByte struct {
	opcode chip8.Opcode
}

// RndVxByte - Cxkk - Set Vx = random byte And kk
func RndVxByte(opcode chip8.Opcode) Instruction {
	return &rndVxByte{opcode: opcode}
}

func (i *rndVxByte) Execute(c *chip8.Chip8) error {
	random := rand.Intn(256)
	c.VX[i.opcode.Vx] = uint8(random) & uint8(i.opcode.KK)

	c.ProgramCounter += 2

	return nil
}

func (i *rndVxByte) String() string {
	return fmt.Sprintf("RND Vx=0x%x byte=0x%x", i.opcode.Vx, i.opcode.KK)
}
