package main

import (
	"github.com/autovelop/playthos"
	// "github.com/autovelop/playthos/collision"
	_ "github.com/autovelop/playthos/glfw"
	_ "github.com/autovelop/playthos/opengl"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
)

func createStartZone(e *engine.Entity, p *std.Vector3) {
	transform := std.NewTransform()
	transform.Set(
		p,
		&std.Vector3{0, 0, 0},
		&std.Vector3{24, 192, 1})
	e.AddComponent(transform)

	quad := render.NewMesh()
	quad.Set(std.QuadMesh)
	e.AddComponent(quad)

	material := render.NewMaterial()
	material.SetColor(
		&std.Color{1.0, 1.0, 1.0, 1.0},
	)

	sa := render.NewImage()
	sa.LoadImage("assets/scene_start.png")
	s := render.NewTexture(sa)
	s.SetSize(24, 192)
	material.SetTexture(s)

	// e.AddComponent(material)
	e.AddComponent(material)
}

func createSafeZone(e *engine.Entity, p *std.Vector3) {
	e.SetTag(1)
	transform := std.NewTransform()
	transform.Set(
		p,
		&std.Vector3{0, 0, 0},
		&std.Vector3{24, 192, 1})
	e.AddComponent(transform)

	quad := render.NewMesh()
	quad.Set(std.QuadMesh)
	e.AddComponent(quad)

	material := render.NewMaterial()
	material.SetColor(
		&std.Color{1.0, 1.0, 1.0, 1.0},
	)
	sa := render.NewImage()
	sa.LoadImage("assets/scene_safe.png")
	s := render.NewTexture(sa)
	s.SetSize(24, 192)
	material.SetTexture(s)

	e.AddComponent(material)

	// collider := collision.NewCollider()
	// collider.Set(transform, &std.Rect{&std.Vector2{0, -96}, 24, 192})
	// e.AddComponent(collider)
}

func createNextZone(e *engine.Entity, p *std.Vector3) {
	e.SetTag(3)
	transform := std.NewTransform()
	transform.Set(
		p,
		&std.Vector3{0, 0, 0},
		&std.Vector3{24, 192, 1})
	e.AddComponent(transform)

	quad := render.NewMesh()
	quad.Set(std.QuadMesh)
	e.AddComponent(quad)

	material := render.NewMaterial()
	material.SetColor(
		&std.Color{1, 0, 0, 0.6},
	)

	e.AddComponent(material)

	// collider := collision.NewCollider()
	// collider.Set(transform, &std.Rect{&std.Vector2{0, -96}, 24, 192})
	// e.AddComponent(collider)
}
