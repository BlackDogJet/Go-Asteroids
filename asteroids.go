package main

import (
	rayLib "github.com/gen2brain/raylib-go/raylib"
)

type AsteroidSize int

const (
	Large AsteroidSize = iota
	Medium
	Small
)

type Asteroid struct {
	position     rayLib.Vector2
	speed        rayLib.Vector2
	size         rayLib.Vector2
	asteroidSize AsteroidSize
}

func (asteroid *Asteroid) Draw() {
	destinationTexture := rayLib.Rectangle{
		X:      asteroid.position.X,
		Y:      asteroid.position.Y,
		Width:  asteroid.size.X,
		Height: asteroid.size.Y,
	}

	rayLib.DrawTexturePro(
		tiles,
		asteroidRectangle,
		destinationTexture,
		rayLib.Vector2{
			X: asteroid.size.X / 2,
			Y: asteroid.size.Y / 2,
		},
		0.0,
		rayLib.White,
	)
}

func (asteroid *Asteroid) Update() {
	asteroid.position = rayLib.Vector2Add(asteroid.position, asteroid.speed)
	wrapPosition(&asteroid.position, asteroid.size.X)
}

func createLargeAsteroid() Asteroid {
	var position rayLib.Vector2
	randomEdge := rayLib.GetRandomValue(0, 3)

	random_X := float32(rayLib.GetRandomValue(0, ScreenWidth))
	random_Y := float32(rayLib.GetRandomValue(0, ScreenHeight))

	switch randomEdge {
	case 0:
		position = rayLib.Vector2{X: random_X, Y: +TileSize}
	case 1:
		position = rayLib.Vector2{X: ScreenWidth + TileSize, Y: random_Y}
	case 2:
		position = rayLib.Vector2{X: random_X, Y: ScreenHeight + TileSize}
	case 3:
		position = rayLib.Vector2{X: -TileSize, Y: random_Y}
	}

	speed := rayLib.Vector2{
		X: float32(rayLib.GetRandomValue(-10, 10)) / 10,
		Y: float32(rayLib.GetRandomValue(-10, 10)) / 10,
	}

	return createAsteroid(Large, position, speed)
}

func createAsteroid(asteroidSize AsteroidSize, position, speed rayLib.Vector2) Asteroid {
	var size rayLib.Vector2

	switch asteroidSize {
	case Large:
		size = rayLib.Vector2{X: TileSize * 1.0, Y: TileSize * 1.0}
	case Medium:
		size = rayLib.Vector2{X: TileSize * 0.7, Y: TileSize * 0.7}
	case Small:
		size = rayLib.Vector2{X: TileSize * 0.4, Y: TileSize * 0.4}
	}

	return Asteroid{
		position:     position,
		speed:        speed,
		size:         size,
		asteroidSize: asteroidSize,
	}
}

func splitAsteroid(asteroid Asteroid) {
	var split int
	var newSize AsteroidSize

	switch size := asteroid.asteroidSize; size {
	case Medium:
		{
			newSize = Small
			split = 4
		}

	case Large:
		{
			newSize = Medium
			split = 2

		}

	default:
		{
			return
		}
	}

	for range split {
		angle := float64(rayLib.GetRandomValue(0, 360))
		direction := getDirectionVector(float32(angle))
		speed := rayLib.Vector2Scale(direction, 2.0)
		newAsteroid := createAsteroid(newSize, asteroid.position, speed)
		asteroids = append(asteroids, newAsteroid)
	}
}
