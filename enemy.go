package main

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/animation"
	"github.com/autovelop/playthos/collision"
	"github.com/autovelop/playthos/physics"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/scripting"
	"github.com/autovelop/playthos/std"
)

func createEnemy(e *engine.Entity, p *std.Vector3, patrolRange *std.Vector2, waypoints []*Waypoint) {
	e.SetTag(2)
	transform := std.NewTransform()
	transform.Set(
		p,
		&std.Vector3{0, 0, 0},
		&std.Vector3{11, 17, 1})
	e.AddComponent(transform)

	quad := render.NewMesh()
	quad.Set(std.QuadMesh)
	e.AddComponent(quad)

	material := render.NewMaterial()
	material.SetColor(
		&std.Color{1.0, 1.0, 1.0, 1.0},
	)

	walkingWestAtlas := render.NewImage()
	walkingWestAtlas.LoadImage("assets", "enemy_west.png")
	walkingWestSprite := render.NewSprite(walkingWestAtlas)
	walkingWestSprite.SetSpriteSize(11, 17)
	walkingWestSpriteCoord := std.Vector2{0, 0}
	walkingWestSprite.SetSpriteOffset(&walkingWestSpriteCoord)

	material.SetSprite(walkingWestSprite)
	e.AddComponent(material)

	walkingEastAtlas := render.NewImage()
	walkingEastAtlas.LoadImage("assets", "enemy_east.png")
	walkingEastSprite := render.NewSprite(walkingEastAtlas)
	walkingEastSprite.SetSpriteSize(11, 17)
	walkingEastSpriteCoord := std.Vector2{0, 0}
	walkingEastSprite.SetSpriteOffset(&walkingEastSpriteCoord)

	walkingUpAtlas := render.NewImage()
	walkingUpAtlas.LoadImage("assets", "enemy_up.png")
	walkingUpSprite := render.NewSprite(walkingUpAtlas)
	walkingUpSprite.SetSpriteSize(11, 17)
	walkingUpSpriteCoord := std.Vector2{0, 0}
	walkingUpSprite.SetSpriteOffset(&walkingUpSpriteCoord)

	walkingDownAtlas := render.NewImage()
	walkingDownAtlas.LoadImage("assets", "enemy_down.png")
	walkingDownSprite := render.NewSprite(walkingDownAtlas)
	walkingDownSprite.SetSpriteSize(11, 17)
	walkingDownSpriteCoord := std.Vector2{0, 0}
	walkingDownSprite.SetSpriteOffset(&walkingDownSpriteCoord)

	walkingWestAnimation := animation.NewClip(1, 180, &walkingWestSpriteCoord)
	walkingWestAnimation.AddKeyFrame(0, 29, &std.Vector2{0, 0})
	walkingWestAnimation.AddKeyFrame(30, 29, &std.Vector2{1, 0})
	walkingWestAnimation.AddKeyFrame(60, 29, &std.Vector2{2, 0})
	walkingWestAnimation.AddKeyFrame(90, 29, &std.Vector2{3, 0})
	walkingWestAnimation.AddKeyFrame(120, 29, &std.Vector2{4, 0})
	walkingWestAnimation.AddKeyFrame(150, 29, &std.Vector2{5, 0})
	walkingWestAnimation.AddKeyFrame(180, 29, &std.Vector2{6, 0})
	// walkingWestAnimation.SetAutoplay(false)
	e.AddComponent(walkingWestAnimation)

	walkingEastAnimation := animation.NewClip(1, 180, &walkingEastSpriteCoord)
	walkingEastAnimation.AddKeyFrame(0, 29, &std.Vector2{0, 0})
	walkingEastAnimation.AddKeyFrame(30, 29, &std.Vector2{1, 0})
	walkingEastAnimation.AddKeyFrame(60, 29, &std.Vector2{2, 0})
	walkingEastAnimation.AddKeyFrame(90, 29, &std.Vector2{3, 0})
	walkingEastAnimation.AddKeyFrame(120, 29, &std.Vector2{4, 0})
	walkingEastAnimation.AddKeyFrame(150, 29, &std.Vector2{5, 0})
	walkingEastAnimation.AddKeyFrame(180, 29, &std.Vector2{6, 0})
	walkingEastAnimation.SetAutoplay(false)
	e.AddComponent(walkingEastAnimation)

	walkingUpAnimation := animation.NewClip(1, 180, &walkingUpSpriteCoord)
	walkingUpAnimation.AddKeyFrame(0, 29, &std.Vector2{0, 0})
	walkingUpAnimation.AddKeyFrame(30, 29, &std.Vector2{1, 0})
	walkingUpAnimation.AddKeyFrame(60, 29, &std.Vector2{2, 0})
	walkingUpAnimation.AddKeyFrame(90, 29, &std.Vector2{3, 0})
	walkingUpAnimation.AddKeyFrame(120, 29, &std.Vector2{4, 0})
	walkingUpAnimation.AddKeyFrame(150, 29, &std.Vector2{5, 0})
	walkingUpAnimation.AddKeyFrame(180, 29, &std.Vector2{6, 0})
	walkingUpAnimation.SetAutoplay(false)
	e.AddComponent(walkingUpAnimation)

	walkingDownAnimation := animation.NewClip(1, 180, &walkingDownSpriteCoord)
	walkingDownAnimation.AddKeyFrame(0, 29, &std.Vector2{0, 0})
	walkingDownAnimation.AddKeyFrame(30, 29, &std.Vector2{1, 0})
	walkingDownAnimation.AddKeyFrame(60, 29, &std.Vector2{2, 0})
	walkingDownAnimation.AddKeyFrame(90, 29, &std.Vector2{3, 0})
	walkingDownAnimation.AddKeyFrame(120, 29, &std.Vector2{4, 0})
	walkingDownAnimation.AddKeyFrame(150, 29, &std.Vector2{5, 0})
	walkingDownAnimation.AddKeyFrame(180, 29, &std.Vector2{6, 0})
	walkingDownAnimation.SetAutoplay(false)
	e.AddComponent(walkingDownAnimation)

	collider := collision.NewCollider()
	collider.Set(transform, &std.Rect{&std.Vector2{-5.5, -8.5}, 11, 17})
	e.AddComponent(collider)

	velocity := physics.NewVelocity()
	e.AddComponent(velocity)

	acc := physics.NewAcceleration()
	e.AddComponent(acc)

	// patrol := animation.NewClip(1, 1200, p)
	// patrol.AddKeyFrame(0, 0, &std.Vector3{p.X, p.Y, 0})
	// patrol.AddKeyFrame(300, 0, &std.Vector3{p.X + patrolRange.X, p.Y, 0})
	// patrol.AddKeyFrame(600, 0, &std.Vector3{p.X + patrolRange.X, p.Y + patrolRange.Y, 0})
	// patrol.AddKeyFrame(900, 0, &std.Vector3{p.X + patrolRange.X, p.Y, 0})
	// patrol.AddKeyFrame(1200, 0, &std.Vector3{p.X, p.Y, 0})
	// e.AddComponent(patrol)

	es := &EnemyState{}
	es.AddWaypoints(waypoints)
	es.Set(e)
	es.SetSprites(walkingEastSprite, walkingWestSprite, walkingUpSprite, walkingDownSprite)
	es.SetAnimations(walkingEastAnimation, walkingWestAnimation, walkingUpAnimation, walkingDownAnimation)
	e.AddComponent(es)

	patrol := scripting.NewScript()
	// ticker := 0.0
	patrol.OnUpdate(func() {
		if len(es.Waypoints()) > 0 && es.Patrol() {
			waypoint := es.Waypoints()[es.WaypointIndex()]
			if waypoint.Bearing().X < 0 {
				if p.X <= waypoint.Dest().X {
					es.NextWaypoint()
				}
			} else if waypoint.Bearing().X > 0 {
				if p.X >= waypoint.Dest().X {
					es.NextWaypoint()
				}
			} else if waypoint.Bearing().Y < 0 {
				if p.Y <= waypoint.Dest().Y {
					es.NextWaypoint()
				}
			} else if waypoint.Bearing().Y > 0 {
				if p.Y >= waypoint.Dest().Y {
					es.NextWaypoint()
				}
			}
			velocity.Set(waypoint.Bearing().X, waypoint.Bearing().Y, 0)
		}
	})
	es.StartPatrol()
	e.AddComponent(patrol)
}
