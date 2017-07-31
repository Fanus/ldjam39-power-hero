package main

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
)

func createCamera(e *engine.Entity) {
	t := std.NewTransform()
	t.Set(
		&std.Vector3{0, 0, 3}, // POSITION
		&std.Vector3{0, 0, 0}, // CENTER
		&std.Vector3{0, 1, 0}, // UP
	)
	e.AddComponent(t)

	camera := render.NewCamera()
	cameraSize := float32(4)
	camera.Set(&cameraSize, &std.Color{0.27, 0.39, 0.54, 0})
	camera.SetTransform(t)

	e.AddComponent(camera)
}
