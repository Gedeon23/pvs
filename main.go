package main

import (
	"fmt"
	"image/color"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// milli astronomical unit
	mAU float32 = 1.0
	// astronomical unit
	AU float32 = 1000.0

	// Megameter
	GM float32 = 1.0 / 150.0
)

type Body interface {
	Draw()
}

type CelestialBody struct {
	Position rl.Vector3
	Radius   float32
	Color    color.RGBA
}

func (a CelestialBody) Draw() {
	rl.DrawSphere(a.Position, a.Radius, a.Color)
}

var Earth CelestialBody = CelestialBody{
	Position: rl.Vector3{X: AU, Y: 0, Z: 0},
	Radius:   6.371 * GM,
	Color:    color.RGBA{121, 110, 211, 255},
}

var Moon CelestialBody = CelestialBody{
	Position: rl.Vector3{X: Earth.Position.X, Y: 0, Z: Earth.Radius + 400*GM},
	Radius:   1.737 * GM,
	Color:    color.RGBA{111, 111, 111, 255},
}

var Sun CelestialBody = CelestialBody{
	Position: rl.Vector3{X: 0, Y: 0, Z: 0},
	Radius:   696.340 * GM,
	Color:    color.RGBA{211, 181, 111, 255},
}

func DrawCelestialBody(body CelestialBody) {
	rl.DrawSphere(body.Position, body.Radius, body.Color)
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rl.DisableCursor()

	camera := rl.Camera3D{
		Position:   Moon.Position,
		Target:     Earth.Position,
		Up:         rl.Vector3{X: 0, Y: 1, Z: 0},
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree)
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("fps: "+strconv.FormatInt(int64(rl.GetFPS()), 10), 10, 10, 20, rl.LightGray)

		rl.BeginMode3D(camera)
		DrawCelestialBody(Earth)
		DrawCelestialBody(Sun)
		DrawCelestialBody(Moon)
		rl.EndMode3D()
		rl.EndDrawing()
	}

	fmt.Printf("Moon coords: %f %f %f, rad: %f", Moon.Position.X, Moon.Position.Y, Moon.Position.Z, Moon.Radius)
}
