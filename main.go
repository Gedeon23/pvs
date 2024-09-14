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
	MM float32 = 1.0 / 150_000.0

	BODY_SCALE float32 = 10_000.0
)

type Body interface {
	Draw()
}

type CelestialBody interface {
	Draw()
}

type Star struct {
	Position rl.Vector3
	Radius   float32
	Color    color.RGBA
}

func (s Star) Draw() {
	rl.DrawSphere(s.Position, s.Radius, s.Color)
}

type Planet struct {
	Position rl.Vector3
	Radius   float32
	Color    color.RGBA
}

func (p Planet) Draw() {
	rl.DrawSphere(p.Position, p.Radius, p.Color)
}

type Moon struct {
	Position rl.Vector3
	Radius   float32
	Color    color.RGBA
}

func (m Moon) Draw() {
	rl.DrawSphere(m.Position, m.Radius, m.Color)
}

var Earth Planet = Planet{
	Position: rl.Vector3{X: AU, Y: 0, Z: 0},
	Radius:   6.371 * MM * BODY_SCALE,
	Color:    color.RGBA{121, 110, 211, 255},
}

var TheMoon Moon = Moon{
	Position: rl.Vector3{X: Earth.Position.X, Y: 0, Z: Earth.Radius + 400*MM*BODY_SCALE},
	Radius:   1.737 * MM * BODY_SCALE,
	Color:    color.RGBA{111, 111, 111, 255},
}

var Sun Star = Star{
	Position: rl.Vector3{X: 0, Y: 0, Z: 0},
	Radius:   696.340 * MM * BODY_SCALE,
	Color:    color.RGBA{211, 181, 111, 255},
}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rl.DisableCursor()

	camera := rl.Camera3D{
		Position:   Earth.Position,
		Target:     Sun.Position,
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
		if rl.Vector3DotProduct(Sun.Position, rl.Vector3Normalize(camera.Target)) > 0.95 {
			rl.DrawText("You are looking at the sun", 10, 35, 20, rl.LightGray)
		}
		TheMoon.Draw()
		Sun.Draw()
		Earth.Draw()
		rl.EndMode3D()
		rl.EndDrawing()
	}
	fmt.Printf("%f", rl.Vector3DotProduct(Sun.Position, rl.Vector3Normalize(camera.Target)))
}
