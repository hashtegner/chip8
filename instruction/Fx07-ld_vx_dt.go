package instruction

import "github.com/hashtegner/chip8/chip8"

// LDVxDt - Fx07 - Set Vx = delay time value
func LDVxDt(opcode chip8.Opcode) Instruction {
	return Undefined(opcode)
}
