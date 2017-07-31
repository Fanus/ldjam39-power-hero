package main

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/std"
)

type Level struct {
	completed bool
	cameraPos *std.Vector3

	enemies []*engine.Entity
}

func (l *Level) SetCameraPosition(c *std.Vector3) {
	l.cameraPos = c
}

func (l *Level) Completed() bool {
	return l.completed
}

func (l *Level) SetCompleted(c bool) {
	l.completed = c
	if c {
		for _, e := range l.enemies {
			es := e.Component(&EnemyState{}).(*EnemyState)
			es.StopPatrol()
		}
	}
}

func (l *Level) AddEnemies(e ...*engine.Entity) {
	l.enemies = append(l.enemies, e...)
}

func (l *Level) Load(p *PlayerState) {
	cTransform := p.Camera().Component(&std.Transform{}).(*std.Transform)
	pTransform := p.Player().Component(&std.Transform{}).(*std.Transform)

	pTransform.SetPosition(l.cameraPos.X-120, pTransform.Position().Y, pTransform.Position().Z)

	cTransform.SetPosition(l.cameraPos.X, l.cameraPos.Y, 3)
	cTransform.SetRotation(l.cameraPos.X, l.cameraPos.Y, 0)

	// tMaterial := p.Terrain().Component(&render.Material{}).(*render.Material)
}
