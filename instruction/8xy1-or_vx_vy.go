package instruction

import "github.com/hashtegner/chip8/chip8"

// OrVxVy - 8xy1 - Set Vx = Vx OR Vy
func OrVxVy(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
