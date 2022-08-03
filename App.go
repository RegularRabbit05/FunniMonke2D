package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"math"
)

func main() {

	var bananas []rl.Vector2

	max := 3.14 * 2
	spd := 360.0 * 1.75
	fr := 0.0
	frl := 0

	cam := rl.Camera3D{Position: rl.Vector3{Z: 3}, Up: rl.Vector3{Y: -1}, Fovy: 30, Projection: rl.CameraPerspective}

	rl.InitAudioDevice()
	rl.InitWindow(1200, 800, "")
	rl.SetTargetFPS(60)

	monkey := rl.LoadTexture("./assets/monkey.png")
	banana := rl.LoadTexture("./assets/banana.png")
	sound := rl.LoadSound("./assets/uhuh.mp3")

	for !rl.WindowShouldClose() {

		fr++
		if fr >= spd {
			fr = 1
		}
		frl += 4
		if frl >= 64 {
			frl = 1
		}

		current := max / spd * fr
		cam.Position.X = float32(math.Cos(current)) * 6
		cam.Position.Z = float32(math.Sin(current)) * 6

		if rl.IsKeyPressed('B') {
			bananas = append(bananas, rl.Vector2{X: float32(rl.GetRandomValue(0, 1200)), Y: float32(rl.GetRandomValue(0, 800))})
			rl.PlaySoundMulti(sound)
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for i := 0; i < len(bananas); i++ {
			rl.DrawTexture(banana, int32(bananas[i].X), int32(bananas[i].Y), rl.White)
		}
		rl.BeginMode3D(cam)
		rl.DrawCubeTexture(monkey, rl.Vector3{}, 1, 1, 1, rl.White)
		rl.EndMode3D()
		rl.EndDrawing()
	}

	rl.UnloadSound(sound)
	rl.UnloadTexture(banana)
	rl.UnloadTexture(monkey)

	rl.CloseAudioDevice()
	rl.CloseWindow()
}