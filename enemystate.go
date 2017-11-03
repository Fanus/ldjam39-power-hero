package main

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/animation"
	_ "github.com/autovelop/playthos/glfw"
	_ "github.com/autovelop/playthos/opengl"
	// "github.com/autovelop/playthos/physics"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
)

type EnemyState struct {
	entity *engine.Entity

	engine.Component
	curWaypointIndex int
	patrol           bool
	waypoints        []*Waypoint

	walkingEastSprite *render.Texture
	walkingWestSprite *render.Texture
	walkingDownSprite *render.Texture
	walkingUpSprite   *render.Texture

	walkingEastAnimation *animation.AnimationClip
	walkingWestAnimation *animation.AnimationClip
	walkingUpAnimation   *animation.AnimationClip
	walkingDownAnimation *animation.AnimationClip
}

func (es *EnemyState) AddWaypoints(w []*Waypoint) {
	es.waypoints = append(es.waypoints, w...)
}

func (es *EnemyState) Waypoints() []*Waypoint {
	return es.waypoints
}

func (es *EnemyState) NextWaypoint() {
	if len(es.waypoints)-1 == es.curWaypointIndex {
		es.curWaypointIndex = 0
	} else {
		es.curWaypointIndex++
	}

	bearing := es.waypoints[es.curWaypointIndex].Bearing()
	material := es.entity.Component(&render.Material{}).(*render.Material)
	if bearing.X < 0 {
		sprite, anim := es.WalkingWest()
		material.SetTexture(sprite)
		anim.Start()
	} else if bearing.X > 0 {
		sprite, anim := es.WalkingEast()
		material.SetTexture(sprite)
		anim.Start()
	} else if bearing.Y < 0 {
		sprite, anim := es.WalkingDown()
		material.SetTexture(sprite)
		anim.Start()
	} else if bearing.Y > 0 {
		sprite, anim := es.WalkingUp()
		material.SetTexture(sprite)
		anim.Start()
	}
}

func (es *EnemyState) WaypointIndex() int {
	return es.curWaypointIndex
}

func (es *EnemyState) Patrol() bool {
	return es.patrol
}

func (es *EnemyState) StopPatrol() {
	es.patrol = false

	bearing := es.waypoints[es.curWaypointIndex].Bearing()
	material := es.entity.Component(&render.Material{}).(*render.Material)
	// velocity := es.entity.Component(&physics.Velocity{}).(*physics.Velocity)
	if bearing.X < 0 {
		sprite, anim := es.WalkingWest()
		material.SetTexture(sprite)
		anim.Stop()
	} else if bearing.X > 0 {
		sprite, anim := es.WalkingEast()
		material.SetTexture(sprite)
		anim.Stop()
	} else if bearing.Y < 0 {
		sprite, anim := es.WalkingDown()
		material.SetTexture(sprite)
		anim.Stop()
	} else if bearing.Y > 0 {
		sprite, anim := es.WalkingUp()
		material.SetTexture(sprite)
		anim.Stop()
	}
	// velocity.SetX(0)
	// velocity.SetY(0)
}

func (es *EnemyState) StartPatrol() {
	es.patrol = true
	es.curWaypointIndex = 0

	if len(es.Waypoints()) > 0 {
		bearing := es.waypoints[es.curWaypointIndex].Bearing()
		material := es.entity.Component(&render.Material{}).(*render.Material)
		if bearing.X < 0 {
			sprite, anim := es.WalkingWest()
			material.SetTexture(sprite)
			anim.Start()
		} else if bearing.X > 0 {
			sprite, anim := es.WalkingEast()
			material.SetTexture(sprite)
			anim.Start()
		} else if bearing.Y < 0 {
			sprite, anim := es.WalkingDown()
			material.SetTexture(sprite)
			anim.Start()
		} else if bearing.Y > 0 {
			sprite, anim := es.WalkingUp()
			material.SetTexture(sprite)
			anim.Start()
		}
	}
}

func (es *EnemyState) WalkingWest() (*render.Texture, *animation.AnimationClip) {
	return es.walkingWestSprite, es.walkingWestAnimation
}

func (es *EnemyState) WalkingEast() (*render.Texture, *animation.AnimationClip) {
	return es.walkingEastSprite, es.walkingEastAnimation
}

func (es *EnemyState) WalkingUp() (*render.Texture, *animation.AnimationClip) {
	return es.walkingUpSprite, es.walkingUpAnimation
}

func (es *EnemyState) WalkingDown() (*render.Texture, *animation.AnimationClip) {
	return es.walkingDownSprite, es.walkingDownAnimation
}
func (es *EnemyState) Set(ent *engine.Entity) {
	es.entity = ent
}

func (es *EnemyState) SetSprites(e *render.Texture, w *render.Texture, u *render.Texture, d *render.Texture) {
	es.walkingEastSprite = e
	es.walkingWestSprite = w
	es.walkingUpSprite = u
	es.walkingDownSprite = d
}

func (es *EnemyState) SetAnimations(e *animation.AnimationClip, w *animation.AnimationClip, u *animation.AnimationClip, d *animation.AnimationClip) {
	es.walkingEastAnimation = e
	es.walkingWestAnimation = w
	es.walkingUpAnimation = u
	es.walkingDownAnimation = d
}

type Waypoint struct {
	bearing *std.Vector3
	dest    *std.Vector3
}

func (w *Waypoint) Bearing() *std.Vector3 {
	return w.bearing
}

func (w *Waypoint) Dest() *std.Vector3 {
	return w.dest
}
