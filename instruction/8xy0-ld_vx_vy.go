package instruction

import "github.com/hashtegner/chip8/chip8"

// LDVxVy - 8xy0 - Set Vx = Vy
func LDVxVy(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
