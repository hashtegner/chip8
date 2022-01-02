package instruction

import (
	"github.com/hashtegner/chip8/chip8"
)

type Instruction interface {
	Execute(*chip8.Chip8) error

	String() string
}
