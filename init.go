package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

var (
	shots              []Shot
	tiles              rayLib.Texture2D
	paused             bool
	player             Player
	victory            bool
	gameOver           bool
	asteroids          []Asteroid
	background         rayLib.Texture2D
	boostRectangle     rayLib.Rectangle
	spriteRectangle    rayLib.Rectangle
	asteroidRectangle  rayLib.Rectangle
	asteroidsDestroyed int
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

	boostRectangle = rayLib.Rectangle{
		X:      TileSize * 7,
		Y:      TileSize * 5,
		Width:  TileSize,
		Height: TileSize,
	}

	asteroidRectangle = rayLib.Rectangle{
		X:      TileSize * 1,
		Y:      TileSize * 4,
		Width:  TileSize,
		Height: TileSize,
	}

	shots = make([]Shot, MaxShots)

	initGame()
}

func initGame() {
	paused = false
	victory = false
	gameOver = false
	asteroidsDestroyed = 0

	asteroids = nil
	for range StartingAsteroids {
		asteroids = append(asteroids, createLargeAsteroid())
	}

	for shot := range shots {
		shots[shot].active = false
	}

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
