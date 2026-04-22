package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

var (
	textBackground rayLib.Texture2D
)

func init() {
	rayLib.InitWindow(ScreenWidth, ScreenHeight, "Go-Asteroids")
	rayLib.SetTargetFPS(60)

	textBackground = rayLib.LoadTexture("assets/space_background.png")
}

func deInit() {
	rayLib.CloseWindow()
	rayLib.UnloadTexture(textBackground)
}
