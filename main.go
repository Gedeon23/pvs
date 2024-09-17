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

	BODY_SCALE  float32 = 10_000.0
	FOCUS_SCALE float32 = 100.0
)

type CelestialBody interface {
	Draw(scale float32)
	GetPosition() rl.Vector3
	GetName() string
}

type Star struct {
	Name     string
	Position rl.Vector3
	Radius   float32
	Color    color.RGBA
}

func (s Star) Draw(scale float32) {
	rl.DrawSphere(s.Position, s.Radius*scale, s.Color)
}

func (s Star) GetPosition() rl.Vector3 {
	return s.Position
}

func (s Star) GetName() string {
	return s.Name
}

type Planet struct {
	Name     string
	Position rl.Vector3
	Radius   float32
	Color    color.RGBA
}

func (p Planet) Draw(scale float32) {
	rl.DrawSphere(p.Position, p.Radius*scale, p.Color)
}

func (p Planet) GetPosition() rl.Vector3 {
	return p.Position
}

func (p Planet) GetName() string {
	return p.Name
}

type Moon struct {
	Name     string
	Position rl.Vector3
	Radius   float32
	Color    color.RGBA
}

func (m Moon) Draw(scale float32) {
	rl.DrawSphere(m.Position, m.Radius*scale, m.Color)
}

func (m Moon) GetPosition() rl.Vector3 {
	return m.Position
}

func (m Moon) GetName() string {
	return m.Name
}

var Earth Planet = Planet{
	Name:     "Earth",
	Position: rl.Vector3{X: AU, Y: 0, Z: 0},
	Radius:   6.371 * MM * BODY_SCALE,
	Color:    color.RGBA{121, 110, 211, 255},
}

var TheMoon Moon = Moon{
	Name: "Moon",
	Position: rl.Vector3{
		X: Earth.Position.X,
		Y: 0,
		Z: Earth.Radius + 400*MM*BODY_SCALE,
	},
	Radius: 1.737 * MM * BODY_SCALE,
	Color:  color.RGBA{111, 111, 111, 255},
}

var Sun Star = Star{
	Name:     "Sun",
	Position: rl.Vector3{X: 0, Y: 0, Z: 0},
	Radius:   696.340 * MM * BODY_SCALE,
	Color:    color.RGBA{211, 181, 111, 255},
}

var CURRENTLY_VIEWED CelestialBody = Sun

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

	celestBodies := []CelestialBody{
		Sun,
		Earth,
		TheMoon,
	}

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree)
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		// Draw Celestial Bodies
		rl.BeginMode3D(camera)
		for _, body := range celestBodies {
			if body == CURRENTLY_VIEWED {
				body.Draw(2.0)
			} else {
				body.Draw(1.0)
			}
		}
		rl.EndMode3D()

		// Overlay Text
		rl.DrawText("fps: "+strconv.FormatInt(int64(rl.GetFPS()), 10), 10, 10, 20, rl.LightGray)

		writtenLines := 0
		CURRENTLY_VIEWED = nil
		for _, body := range celestBodies {
			angle := rl.Vector3Angle(rl.Vector3Subtract(body.GetPosition(), camera.Position), rl.Vector3Subtract(camera.Target, camera.Position))
			if angle < 0.1 {
				rl.DrawText("You are looking at the "+body.GetName()+fmt.Sprintf(" Angle: %f", angle), 10, int32(35+writtenLines*25), 20, rl.LightGray)
				CURRENTLY_VIEWED = body
				writtenLines++
			}
		}
		rl.EndDrawing()
	}
}
