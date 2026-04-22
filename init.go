package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

func init() {
	rayLib.InitWindow(ScreenWidth, ScreenHeight, "Go-Asteroids")
	rayLib.SetTargetFPS(60)
}

func deInit() {
	rayLib.CloseWindow()
}
