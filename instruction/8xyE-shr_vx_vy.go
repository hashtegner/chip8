package instruction

import "github.com/hashtegner/chip8/chip8"

// SHRVxVyMostSignificant - 8xyE - Set Vx = Vx SHR 1 if the most-significant bit of Vx is 1.
func SHRVxVyMostSignificant(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
