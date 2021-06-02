package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
)

type mrowisko struct {
	mrowki []mrowka
	x,y int32 // punkt startowy
}

type mrowka struct {
	x, y int32
}

func StworzMrowisko(x,y int32, ile_mrowek int) (mrowisko) {
	mrowki := make([]mrowka, ile_mrowek)

	m := mrowisko{x: x, y: y, mrowki: mrowki}

	return m
}

func WyczyscEkran() {

}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("mrowisko", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 800, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)
	window.UpdateSurface()

	M := StworzMrowisko(400, 400, 10)
	_ = M

	rM := sdl.Rect{M.x, M.y, 1, 1}
	surface.FillRect(&rM, 0x00ffff00)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}

			// fmt.Println("dziala")

			// time.Sleep(time.Second / 2)
			sdl.Delay(100)
		}
	}
}