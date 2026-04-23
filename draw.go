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
		Width:  float32(background.Width),
		Height: float32(background.Height),
	}

	backgroundDestinationRec := rayLib.Rectangle{
		X:      0,
		Y:      0,
		Width:  float32(ScreenWidth),
		Height: float32(ScreenHeight),
	}

	rayLib.DrawTexturePro(
		background,
		backgroundSourceRec,
		backgroundDestinationRec,
		rayLib.Vector2{X: 0, Y: 0}, 0, rayLib.White)

	// Draw player
	player.Draw()

	// Draw asteroids
	for asteroid := range asteroids {
		asteroids[asteroid].Draw()
	}

	if gameOver {
		drawCenteredText("Game over", ScreenHeight/2, 50, rayLib.Red)
	}

	rayLib.ClearBackground(rayLib.Black)
	rayLib.DrawText("Score: 0", 10, 10, 20, rayLib.Gray)
	rayLib.EndDrawing()
}

func update() {
	if !gameOver {
		player.Update()

		for asteroid := range asteroids {
			asteroids[asteroid].Update()
		}

		checkCollisions()
	}
}

func drawCenteredText(text string, position_y int32, fontSize int32, color rayLib.Color) {
	textWidth := rayLib.MeasureText(text, fontSize)
	rayLib.DrawText(text, (ScreenWidth/2)-(textWidth/2), position_y, fontSize, color)
}
