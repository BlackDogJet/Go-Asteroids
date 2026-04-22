package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	defer deInit()

	for !rayLib.WindowShouldClose() {
		draw()
	}
}
