package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

var (
	tiles             rayLib.Texture2D
	player            Player
	asteroids         []Asteroid
	background        rayLib.Texture2D
	boostRectangle    rayLib.Rectangle
	spriteRectangle   rayLib.Rectangle
	asteroidRectangle rayLib.Rectangle
)
