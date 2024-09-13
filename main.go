package main

import (
	"image/color"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var AU float32 = 149_597_870
var EARTH_POS rl.Vector3 = rl.Vector3{X: 0, Y: 0, Z: 0}

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rl.DisableCursor()

	camera := rl.Camera3D{
		Position:   rl.Vector3{X: 10, Y: 10, Z: 10},
		Target:     rl.Vector3{X: 0, Y: 0, Z: 0},
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
		rl.DrawSphere(EARTH_POS, 1.0, color.RGBA{121, 110, 211, 255})
		rl.EndMode3D()

		rl.EndDrawing()
	}
}
