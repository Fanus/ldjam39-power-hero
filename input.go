package main

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/keyboard"
	"github.com/autovelop/playthos/physics"
	"github.com/autovelop/playthos/render"
)

func handlePlayerInput(kb engine.Listener, e *engine.Entity) {
	ps := e.Component(&PlayerState{}).(*PlayerState)
	keysDown := 0
	kb.On(keyboard.KeyUp, func(action ...uint) {
		velocity := e.Component(&physics.Velocity{}).(*physics.Velocity)
		switch action[0] {
		case keyboard.ActionPress:
			keysDown++
			if ps.Running() {
				velocity.SetX(0)
			}
			velocity.SetY(0.2)

			material := e.Component(&render.Material{}).(*render.Material)
			sprite, anim := ps.WalkingUp()
			material.SetSprite(sprite)
			anim.Start()

			ps.SetVertical(true)
			break
		case keyboard.ActionRelease:
			if !ps.Running() {
				velocity.SetY(0)

				material := e.Component(&render.Material{}).(*render.Material)
				sprite, anim := ps.Idle()
				material.SetSprite(sprite)
				anim.Start()
			}
			keysDown--
			break
		}
	})
	kb.On(keyboard.KeyDown, func(action ...uint) {
		velocity := e.Component(&physics.Velocity{}).(*physics.Velocity)
		switch action[0] {
		case keyboard.ActionPress:
			keysDown++
			if ps.Running() {
				velocity.SetX(0)
			}
			velocity.SetY(-0.2)

			material := e.Component(&render.Material{}).(*render.Material)
			sprite, anim := ps.WalkingDown()
			material.SetSprite(sprite)
			anim.Start()

			ps.SetVertical(true)
			break
		case keyboard.ActionRelease:
			if !ps.Running() {
				velocity.SetY(0)

				material := e.Component(&render.Material{}).(*render.Material)
				sprite, anim := ps.Idle()
				material.SetSprite(sprite)
				anim.Start()
			}
			keysDown--
			break
		}
	})
	kb.On(keyboard.KeyRight, func(action ...uint) {
		velocity := e.Component(&physics.Velocity{}).(*physics.Velocity)
		switch action[0] {
		case keyboard.ActionPress:
			if ps.Completed() {
				return
			}
			keysDown++
			if !ps.Running() {
				ps.SetRunning(true)
			} else {
				velocity.SetY(0)
			}
			velocity.SetX(0.2)

			material := e.Component(&render.Material{}).(*render.Material)
			sprite, anim := ps.WalkingW()
			material.SetSprite(sprite)
			anim.Start()

			ps.SetVertical(false)
			break
		case keyboard.ActionRelease:
			if !ps.Running() {
				velocity.SetX(0)
			}
			// material.SetSprite(ps.Idle())
			keysDown--
			break
		}
	})
}
