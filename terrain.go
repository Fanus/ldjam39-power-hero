package main

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
)

func createTerrain(e *engine.Entity, p *std.Vector3) {
	transform := std.NewTransform()
	transform.Set(
		p,
		&std.Vector3{0, 0, 0},
		&std.Vector3{208, 192, 1})
	e.AddComponent(transform)

	quad := render.NewMesh()
	quad.Set(std.QuadMesh)
	e.AddComponent(quad)

	material := render.NewMaterial()
	material.SetColor(
		&std.Color{1.0, 1.0, 1.0, 1.0},
	)

	sa := render.NewImage()
	sa.LoadImage("assets/scene_ground_1.png")
	s := render.NewTexture(sa)
	s.SetSize(208, 192)
	material.SetTexture(s)

	e.AddComponent(material)
}
