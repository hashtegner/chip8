package instruction

import "github.com/hashtegner/chip8/chip8"

// SENVXByte - 4xkk - Skip next instruction if Vx != kk
func SENVXByte(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
