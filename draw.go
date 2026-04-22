package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

func draw() {
	rayLib.BeginDrawing()

	// Set background
	backgroundSourceRec := rayLib.Rectangle{
		X:      0,
		Y:      0,
		Width:  float32(textBackground.Width),
		Height: float32(textBackground.Height),
	}

	backgroundDestinationRec := rayLib.Rectangle{
		X:      0,
		Y:      0,
		Width:  float32(ScreenWidth),
		Height: float32(ScreenHeight),
	}

	rayLib.DrawTexturePro(
		textBackground,
		backgroundSourceRec,
		backgroundDestinationRec,
		rayLib.Vector2{X: 0, Y: 0}, 0, rayLib.White)

	rayLib.ClearBackground(rayLib.Black)
	rayLib.DrawText("Score: 0", 10, 10, 20, rayLib.Gray)
	rayLib.EndDrawing()
}
