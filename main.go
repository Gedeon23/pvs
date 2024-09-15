package main

import (
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
	GetPosition() rl.Vector3
	GetName() string
}

type Star struct {
	Name     string
	Position rl.Vector3
	Radius   float32
	Color    color.RGBA
}

func (s Star) Draw() {
	rl.DrawSphere(s.Position, s.Radius, s.Color)
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

func (p Planet) Draw() {
	rl.DrawSphere(p.Position, p.Radius, p.Color)
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

func (m Moon) Draw() {
	rl.DrawSphere(m.Position, m.Radius, m.Color)
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
	Name:     "Moon",
	Position: rl.Vector3{X: Earth.Position.X, Y: 0, Z: Earth.Radius + 400*MM*BODY_SCALE},
	Radius:   1.737 * MM * BODY_SCALE,
	Color:    color.RGBA{111, 111, 111, 255},
}

var Sun Star = Star{
	Name:     "Sun",
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

	celestBodies := []CelestialBody{
		Sun,
		Earth,
		TheMoon,
	}

	for !rl.WindowShouldClose() {
		rl.UpdateCamera(&camera, rl.CameraFree)
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("fps: "+strconv.FormatInt(int64(rl.GetFPS()), 10), 10, 10, 20, rl.LightGray)

		writtenLines := 0
		for _, body := range celestBodies {
			sun_camera_proj := rl.Vector3DotProduct(rl.Vector3Subtract(body.GetPosition(), camera.Position), rl.Vector3Normalize(camera.Target))
			if sun_camera_proj > 0.95 {
				rl.DrawText("You are looking at the "+body.GetName(), 10, int32(35+writtenLines*25), 20, rl.LightGray)
				writtenLines++
			}
		}

		rl.BeginMode3D(camera)
		for _, body := range celestBodies {
			body.Draw()
		}
		rl.EndMode3D()
		rl.EndDrawing()
	}
}
