package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

type Shot struct {
	position rayLib.Vector2
	speed    rayLib.Vector2
	radius   float32
	active   bool
}

func (shot *Shot) Draw() {
	if shot.active {
		rayLib.DrawCircleV(shot.position, shot.radius, rayLib.Yellow)
	}
}

func (shot *Shot) Update() {
	if shot.active {
		shot.position.X += shot.speed.X
		shot.position.Y -= shot.speed.Y
		if shot.position.X < 0 || shot.position.X > ScreenWidth ||
			shot.position.Y < 0 || shot.position.Y > ScreenHeight {
			shot.active = false
		}
	}
}

func fireShot() {
	for shot := range shots {
		if !shots[shot].active {
			shots[shot].position = player.position
			shots[shot].active = true

			shotDirection := getDirectionVector(player.rotation)
			shotVelocity := rayLib.Vector2Scale(shotDirection, ShotSpeed)
			playerVelocity := rayLib.Vector2Scale(player.speed, player.acceleration)
			shots[shot].speed = rayLib.Vector2Add(playerVelocity, shotVelocity)
			shots[shot].radius = 2
			break
		}
	}
}
