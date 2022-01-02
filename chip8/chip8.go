package chip8

const (
	BufferSize          = 0x200 // 512
	ScreenWidth         = 64
	ScreenHeight        = 32
	MemorySize          = 4096
	AvailableMemorySize = MemorySize - BufferSize
	FontSetSize         = 80
)

type ScreenBuffer = [ScreenHeight][ScreenWidth]uint8

var FontSet = [FontSetSize]uint8{
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

type Chip8 struct {
	Memory [MemorySize]uint8 // 4k memory
	VX     [16]uint8         // cpu registers: v0-v15

	Stack        [16]uint16 // stack and stack pointer
	StackPointer uint16

	Keys [16]uint8 // keypad

	Screen ScreenBuffer // graphics
	Draw   bool         // draw flag

	Index          uint16 // index register
	ProgramCounter uint16 // program counter

	DelayTimer uint8 // timer counters
	SoundTimer uint8
}