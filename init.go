package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

func init() {
	rayLib.InitWindow(ScreenWidth, ScreenHeight, "Go-Asteroids")
	rayLib.SetTargetFPS(60)

	tiles = rayLib.LoadTexture("assets/tilesheet.png")
	background = rayLib.LoadTexture("assets/space_background.png")

	spriteRectangle = rayLib.Rectangle{
		X:      0,
		Y:      TileSize * 2,
		Width:  TileSize,
		Height: TileSize,
	}

	initGame()
}

func initGame() {
	player = Player{
		position:     rayLib.Vector2{X: 400, Y: 200},
		speed:        rayLib.Vector2{X: 0.0, Y: 0.0},
		size:         rayLib.Vector2{X: TileSize, Y: TileSize},
		rotation:     0.0,
		acceleration: 0.0,
		isBoosting:   false,
	}
}

func deInit() {
	rayLib.CloseWindow()
	rayLib.UnloadTexture(tiles)
	rayLib.UnloadTexture(background)
}
