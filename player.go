package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	position     rayLib.Vector2
	speed        rayLib.Vector2
	size         rayLib.Vector2
	acceleration float32
	rotation     float32
	isBoosting   bool
}

func (player *Player) Draw() {
	destinationTexture := rayLib.Rectangle{
		X:      player.position.X,
		Y:      player.position.Y,
		Width:  player.size.X,
		Height: player.size.Y,
	}

	rayLib.DrawTexturePro(
		tiles,
		spriteRectangle,
		destinationTexture,
		rayLib.Vector2{
			X: player.size.X / 2,
			Y: player.size.Y / 2,
		},
		player.rotation,
		rayLib.White,
	)
}
