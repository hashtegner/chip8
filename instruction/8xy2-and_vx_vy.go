package instruction

import "github.com/hashtegner/chip8/chip8"

// AndVxVy - 8xy2 - Set Vx = Vx AND Vy
func AndVxVy(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
