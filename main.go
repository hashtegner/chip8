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
	rand.Seed(time.Now().UTC().Unix())

	cartridge, err := cartridge.New("roms/chip8-picture.ch8")
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

	clockRate := time.Duration(16 * time.Millisecond)
	clock := time.Now()

loop:
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				print("quit")
				break loop
			}
		}

		err := emulation.Cycle()
		if err != nil {
			log.Printf("error on execute cycle: %v\n", err)
		}

		if emulation.ShouldDraw() {
			graphics.Draw(emulation.ScreenBuffer())
		}

		now := time.Now()

		if now.Sub(clock) > clockRate {
			emulation.Tick()
			clock = now
		}
	}
}
