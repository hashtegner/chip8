package instruction

import "github.com/hashtegner/chip8/chip8"

// RndVxByte - Cxkk - Set Vx = random byte And kk
func RndVxByte(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
