package instruction

import "github.com/hashtegner/chip8/chip8"

// XorVxVy - 8xy3 - Set Vx = Vx XOR Vy
func XorVxVy(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
