package instruction

import "github.com/hashtegner/chip8/chip8"

// SubnVxVy - 8xy7 - Set Vx = Vy - Vx, set VF = Not borrow.
func SubnVxVy(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
