package instruction

import "github.com/hashtegner/chip8/chip8"

// SNEVxVy - 9xy0 - Skip next instruction if Vx != Vy
func SNEVxVy(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
