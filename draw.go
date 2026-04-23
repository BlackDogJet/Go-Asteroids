package main

import (
	"fmt"

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

	for shot := range shots {
		shots[shot].Draw()
	}

	if gameOver {
		drawCenteredText("Game over", ScreenHeight/2, 50, rayLib.Red)
		drawCenteredText("Press R to restart", ScreenHeight/2+60, 20, rayLib.White)
	}

	if victory {
		drawCenteredText("Congrats, you won!", ScreenHeight/2, 50, rayLib.Green)
		drawCenteredText("Press R to restart", ScreenHeight/2+60, 20, rayLib.White)
	}

	if paused {
		drawCenteredText("Game Paused", ScreenHeight/2, 50, rayLib.Green)
		drawCenteredText("Press P to resume", ScreenHeight/2+60, 20, rayLib.White)
	}

	rayLib.ClearBackground(rayLib.Black)
	rayLib.DrawText(fmt.Sprintf("Score %d", asteroidsDestroyed), 10, 10, 20, rayLib.Gray)

	// Pause
	pauseTextSize := rayLib.MeasureText("[P]ause", 20)
	rayLib.DrawText("[P]ause", ScreenWidth-pauseTextSize-10, 10, 20, rayLib.Gray)

	rayLib.EndDrawing()
}

func update() {
	if len(asteroids) == 0 {
		victory = true
	}

	if rayLib.IsKeyDown(rayLib.KeyP) {
		paused = !paused
	}

	if (gameOver || victory) && rayLib.IsKeyDown(rayLib.KeyR) {
		initGame()
	}

	if !gameOver && !paused && !victory {
		player.Update()

		for asteroid := range asteroids {
			asteroids[asteroid].Update()
		}

		for shot := range shots {
			shots[shot].Update()
		}

		checkCollisions()
	}
}

func drawCenteredText(text string, position_y int32, fontSize int32, color rayLib.Color) {
	textWidth := rayLib.MeasureText(text, fontSize)
	rayLib.DrawText(text, (ScreenWidth/2)-(textWidth/2), position_y, fontSize, color)
}

func checkCollisions() {
	for asteroid := len(asteroids) - 1; asteroid >= 0; asteroid-- {
		if rayLib.CheckCollisionCircles(
			player.position,
			player.size.X/4,
			asteroids[asteroid].position,
			asteroids[asteroid].size.X/4,
		) {
			gameOver = true
		}

		for shot := range shots {
			if shots[shot].active {
				if rayLib.CheckCollisionCircles(
					shots[shot].position,
					shots[shot].radius,
					asteroids[asteroid].position,
					asteroids[asteroid].size.X/2,
				) {
					shots[shot].active = false
					splitAsteroid(asteroids[asteroid])
					asteroids = append(asteroids[:asteroid], asteroids[asteroid+1:]...)
					asteroidsDestroyed++
					break
				}
			}
		}
	}
}
