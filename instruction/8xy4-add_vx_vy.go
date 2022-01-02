package instruction

import "github.com/hashtegner/chip8/chip8"

// AddVxVy - 8xy4 - Set Vx = Vx + Vy, set VF = carry
func AddVxVy(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
