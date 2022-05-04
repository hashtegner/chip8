package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hashtegner/chip8/cartridge"
	"github.com/hashtegner/chip8/chip8"
	"github.com/hashtegner/chip8/emulation"
	"github.com/hashtegner/chip8/graphics"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("You need to inform rom file")
	}

	rom := os.Args[1]

	rand.Seed(time.Now().UTC().Unix())

	cartridge, err := cartridge.New(rom)
	if err != nil {
		log.Fatal(err)
	}

	graphics, err := graphics.New("Chip8", chip8.ScreenWidth, chip8.ScreenHeight)
	if err != nil {
		log.Fatal(err)
	}
	defer graphics.Close()

	emulation, err := emulation.New(cartridge)
	if err != nil {
		log.Fatal(err)
	}

loop:
	for {
		start := time.Now()
		err := emulation.Cycle()
		if err != nil {
			log.Printf("error on execute cycle: %v\n", err)
		}

		if emulation.ShouldDraw() {
			graphics.Draw(emulation.ScreenBuffer())
		}

		emulation.Tick()

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch evt := event.(type) {
			case *sdl.QuitEvent:
				print("quit")
				break loop
			case *sdl.KeyboardEvent:
				down := evt.Type == sdl.KEYDOWN

				switch evt.Keysym.Sym {
				case sdl.K_1:
					emulation.Key(0x1, down)
				case sdl.K_2:
					emulation.Key(0x2, down)
				case sdl.K_3:
					emulation.Key(0x3, down)
				case sdl.K_4:
					emulation.Key(0xC, down)
				case sdl.K_q:
					emulation.Key(0x4, down)
				case sdl.K_w:
					emulation.Key(0x5, down)
				case sdl.K_e:
					emulation.Key(0x6, down)
				case sdl.K_r:
					emulation.Key(0xD, down)
				case sdl.K_a:
					emulation.Key(0x7, down)
				case sdl.K_s:
					emulation.Key(0x8, down)
				case sdl.K_d:
					emulation.Key(0x9, down)
				case sdl.K_f:
					emulation.Key(0xE, down)
				case sdl.K_z:
					emulation.Key(0xA, down)
				case sdl.K_x:
					emulation.Key(0x0, down)
				case sdl.K_c:
					emulation.Key(0xB, down)
				case sdl.K_v:
					emulation.Key(0xF, down)
				}
			}
		}

		sdl.Delay(500 / 30)

		fps := (1 * time.Second) / time.Since(start)
		graphics.DisplayFPS(fps.Nanoseconds())
	}
}
