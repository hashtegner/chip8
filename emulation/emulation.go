package emulation

import (
	"fmt"

	"github.com/hashtegner/chip8/cartridge"
	"github.com/hashtegner/chip8/chip8"
	"github.com/hashtegner/chip8/opcode"
	"github.com/pkg/errors"
)

type Emulation struct {
	chip8 *chip8.Chip8
}

func New(cartridge *cartridge.Cartridge) (*Emulation, error) {
	chip8 := &chip8.Chip8{
		ProgramCounter: chip8.BufferSize,
		Draw:           true,
	}

	initializeMemory(chip8)
	err := loadProgram(chip8, cartridge)
	if err != nil {
		return nil, errors.Wrap(err, "error on load cartridge")
	}

	return &Emulation{chip8}, nil
}

func (e *Emulation) ShouldSound() bool {
	return e.chip8.SoundTimer > 0
}

func (e *Emulation) ScreenBuffer() chip8.ScreenBuffer {
	e.chip8.Draw = false

	return e.chip8.Screen
}

func (e *Emulation) ShouldDraw() bool {
	return e.chip8.Draw
}

func (e *Emulation) Cycle() error {
	op := opcode.Fetch(e.chip8)
	inst := opcode.Decode(op)

	return inst.Execute(e.chip8)
}

// initializeMemory load fontset into chip8 memory
func initializeMemory(c *chip8.Chip8) {
	for i := 0; i < chip8.FontSetSize; i++ {
		c.Memory[i] = chip8.FontSet[i]
	}
}

// loadProgram fills chip8 memory with program
func loadProgram(c *chip8.Chip8, cr *cartridge.Cartridge) error {
	size := cr.Size()
	if cr.Size() > chip8.AvailableMemorySize {
		return fmt.Errorf("program size (%v) is bigger than memory (%v)", size, chip8.AvailableMemorySize)
	}

	buffer := cr.Buffer()

	for i := 0; i < len(buffer); i++ {
		c.Memory[i+chip8.BufferSize] = buffer[i]
	}

	return nil
}
