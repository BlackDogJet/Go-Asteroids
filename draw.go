package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

func draw() {
	rayLib.BeginDrawing()
	rayLib.ClearBackground(rayLib.Black)
	rayLib.DrawText("Score: 0", 10, 10, 20, rayLib.Gray)
	rayLib.EndDrawing()
}
