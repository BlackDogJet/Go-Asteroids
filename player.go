package main

import (
	"math"

	rayLib "github.com/gen2brain/raylib-go/raylib"
	rl "github.com/gen2brain/raylib-go/raylib"
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

func (player *Player) Update() {
	// Rotate plaer with arrow keys or WASD
	if rayLib.IsKeyDown(rayLib.KeyLeft) || rayLib.IsKeyDown(rayLib.KeyA) {
		player.rotation -= RotationSpeed
	} else if rayLib.IsKeyDown(rayLib.KeyRight) || rayLib.IsKeyDown(rayLib.KeyD) {
		player.rotation += RotationSpeed
	} else if rayLib.IsKeyDown(rayLib.KeyUp) || rayLib.IsKeyDown(rayLib.KeyW) {
		// Speed the player up
		if player.acceleration < 0.9 {
			player.acceleration += 0.1
		}
	} else if rayLib.IsKeyDown(rayLib.KeyDown) || rayLib.IsKeyDown(rayLib.KeyS) {
		// Slow the player down
		if player.acceleration > 0.0 {
			player.acceleration -= 0.05
		}

		if player.acceleration < 0 {
			player.acceleration = 0
		}
	}

	// Get direction the player is pointing
	direction := getDirectionVector(player.rotation)

	// Move in that direction
	player.speed = rayLib.Vector2Scale(direction, PlayerSpeed)

	// Accelerate in that direction
	player.position.X += player.speed.X * player.acceleration
	player.position.Y -= player.speed.Y * player.acceleration

	// Wrap screen so we don't lose our player
	wrapPosition(&player.position, TileSize)
}

func getDirectionVector(rotation float32) rayLib.Vector2 {
	radians := float64(rotation) * rayLib.Deg2rad

	return rayLib.Vector2{
		X: float32(math.Sin(radians)),
		Y: float32(math.Cos(radians)),
	}
}

func wrapPosition(position *rl.Vector2, objectSize float32) {
	// If we go off the left side of the screen
	if position.X > ScreenWidth+objectSize {
		position.X = -objectSize
	}

	// If we go off the right side of the screen
	if position.X < -objectSize {
		position.X = ScreenWidth + objectSize
	}

	// If we go off the bottom of the screen
	if position.Y > ScreenHeight+objectSize {
		position.Y = -objectSize
	}

	// If we go off the top of the screen
	if position.Y < -objectSize {
		position.Y = ScreenHeight + objectSize
	}
}
