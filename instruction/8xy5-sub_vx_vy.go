package instruction

import "github.com/hashtegner/chip8/chip8"

// SubVxVy - 8xy5 - Set Vx = Vx - Vy, set VF = Not borrow
func SubVxVy(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
