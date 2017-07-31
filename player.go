package main

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/animation"
	"github.com/autovelop/playthos/collision"
	"github.com/autovelop/playthos/physics"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/scripting"
	"github.com/autovelop/playthos/std"
	// "log"
	"math"
)

var player *engine.Entity

func createPlayer(e *engine.Entity, p *std.Vector3, c *engine.Entity, ls []*Level) {
	e.SetTag(99)

	player = e

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
		&std.Color{1, 1, 1, 1},
	)

	idleAtlas := render.NewImage()
	idleAtlas.LoadImage("assets", "player_idle.png")
	idleSprite := render.NewSprite(idleAtlas)
	idleSprite.SetSpriteSize(11, 17)
	idleSpriteCoord := std.Vector2{0, 0}
	idleSprite.SetSpriteOffset(&idleSpriteCoord)

	material.SetSprite(idleSprite)
	e.AddComponent(material)

	walkingEastAtlas := render.NewImage()
	walkingEastAtlas.LoadImage("assets", "player_east.png")
	walkingEastSprite := render.NewSprite(walkingEastAtlas)
	walkingEastSprite.SetSpriteSize(11, 17)
	walkingEastSpriteCoord := std.Vector2{0, 0}
	walkingEastSprite.SetSpriteOffset(&walkingEastSpriteCoord)

	walkingUpAtlas := render.NewImage()
	walkingUpAtlas.LoadImage("assets", "player_up.png")
	walkingUpSprite := render.NewSprite(walkingUpAtlas)
	walkingUpSprite.SetSpriteSize(11, 17)
	walkingUpSpriteCoord := std.Vector2{0, 0}
	walkingUpSprite.SetSpriteOffset(&walkingUpSpriteCoord)

	walkingDownAtlas := render.NewImage()
	walkingDownAtlas.LoadImage("assets", "player_down.png")
	walkingDownSprite := render.NewSprite(walkingDownAtlas)
	walkingDownSprite.SetSpriteSize(11, 17)
	walkingDownSpriteCoord := std.Vector2{0, 0}
	walkingDownSprite.SetSpriteOffset(&walkingDownSpriteCoord)

	idleAnimation := animation.NewClip(1, 300, &idleSpriteCoord)
	idleAnimation.AddKeyFrame(0, 49, &std.Vector2{0, 0})
	idleAnimation.AddKeyFrame(50, 49, &std.Vector2{1, 0})
	idleAnimation.AddKeyFrame(100, 49, &std.Vector2{2, 0})
	idleAnimation.AddKeyFrame(150, 49, &std.Vector2{3, 0})
	idleAnimation.AddKeyFrame(200, 49, &std.Vector2{4, 0})
	idleAnimation.AddKeyFrame(250, 49, &std.Vector2{5, 0})
	idleAnimation.AddKeyFrame(300, 49, &std.Vector2{6, 0})
	// idleAnimation.SetAutoplay(false)
	e.AddComponent(idleAnimation)

	walkingEastAnimation := animation.NewClip(1, 300, &walkingEastSpriteCoord)
	walkingEastAnimation.AddKeyFrame(0, 49, &std.Vector2{0, 0})
	walkingEastAnimation.AddKeyFrame(50, 49, &std.Vector2{1, 0})
	walkingEastAnimation.AddKeyFrame(100, 49, &std.Vector2{2, 0})
	walkingEastAnimation.AddKeyFrame(150, 49, &std.Vector2{3, 0})
	walkingEastAnimation.AddKeyFrame(200, 49, &std.Vector2{4, 0})
	walkingEastAnimation.AddKeyFrame(250, 49, &std.Vector2{5, 0})
	walkingEastAnimation.AddKeyFrame(300, 49, &std.Vector2{6, 0})
	walkingEastAnimation.SetAutoplay(false)
	e.AddComponent(walkingEastAnimation)

	walkingUpAnimation := animation.NewClip(1, 300, &walkingUpSpriteCoord)
	walkingUpAnimation.AddKeyFrame(0, 49, &std.Vector2{0, 0})
	walkingUpAnimation.AddKeyFrame(50, 49, &std.Vector2{1, 0})
	walkingUpAnimation.AddKeyFrame(100, 49, &std.Vector2{2, 0})
	walkingUpAnimation.AddKeyFrame(150, 49, &std.Vector2{3, 0})
	walkingUpAnimation.AddKeyFrame(200, 49, &std.Vector2{4, 0})
	walkingUpAnimation.AddKeyFrame(250, 49, &std.Vector2{5, 0})
	walkingUpAnimation.AddKeyFrame(300, 49, &std.Vector2{6, 0})
	walkingUpAnimation.SetAutoplay(false)
	e.AddComponent(walkingUpAnimation)

	walkingDownAnimation := animation.NewClip(1, 300, &walkingDownSpriteCoord)
	walkingDownAnimation.AddKeyFrame(0, 49, &std.Vector2{0, 0})
	walkingDownAnimation.AddKeyFrame(50, 49, &std.Vector2{1, 0})
	walkingDownAnimation.AddKeyFrame(100, 49, &std.Vector2{2, 0})
	walkingDownAnimation.AddKeyFrame(150, 49, &std.Vector2{3, 0})
	walkingDownAnimation.AddKeyFrame(200, 49, &std.Vector2{4, 0})
	walkingDownAnimation.AddKeyFrame(250, 49, &std.Vector2{5, 0})
	walkingDownAnimation.AddKeyFrame(300, 49, &std.Vector2{6, 0})
	walkingDownAnimation.SetAutoplay(false)
	e.AddComponent(walkingDownAnimation)

	velocity := physics.NewVelocity()
	e.AddComponent(velocity)

	acc := physics.NewAcceleration()
	e.AddComponent(acc)

	ps := &PlayerState{}
	ps.Set([]*engine.Entity{}, &std.Vector3{p.X, p.Y, p.Z}, false, e, c, ls)
	ps.SetSprites(idleSprite, walkingEastSprite, walkingUpSprite, walkingDownSprite)
	ps.SetAnimations(idleAnimation, walkingEastAnimation, walkingUpAnimation, walkingDownAnimation)
	e.AddComponent(ps)

	collider := collision.NewCollider()
	collider.OnHit(func(other *engine.Entity) {
		switch other.Tag() {
		case 1: // zone
			ps.SetRunning(false)
			cTransform := c.Component(&std.Transform{}).(*std.Transform)
			velocity.SetX(0)
			ps.SetLevelCompleted(ps.CurrentLevel())
			other.SetTag(0)
			cPos := cTransform.Position()
			transform.SetPosition(cPos.X+120, transform.Position().Y, transform.Position().Z)

			material := other.Component(&render.Material{}).(*render.Material)
			sa := render.NewImage()
			sa.LoadImage("assets", "scene_safe_connected.png")
			s := render.NewSprite(sa)
			s.SetSpriteSize(24, 192)
			material.SetSprite(s)
			break
		case 2: // enemy
			ps.SetRunning(false)
			velocity.SetX(0)
			velocity.SetY(0)
			cableBlocks := ps.CableBlocks()
			for i := 0; i < len(cableBlocks); i++ {
				e.Engine().DeleteEntity(cableBlocks[i])
			}

			ps.ResetLevel()
			// level.Load(ps)
			// startPos := ps.StartPos()

			// transform.SetPosition(startPos.X, startPos.Y, startPos.Z)
			break
		case 3: // next level
			ps.SetRunning(false)
			velocity.SetX(0)
			velocity.SetY(0)
			ps.NextLevel()
			break
		}
	})
	collider.Set(transform, &std.Rect{&std.Vector2{-5.5, -8.5}, 11, 17})
	e.AddComponent(collider)
}

func createCableBlock(e *engine.Entity, p *std.Vector3, player *engine.Entity, vertical bool) {
	e.SetTag(98)
	transform := std.NewTransform()
	if vertical {
		transform.Set(
			p,
			&std.Vector3{0, 0, 90},
			&std.Vector3{4, 1.5, 1})
	} else {
		transform.Set(
			p,
			&std.Vector3{0, 0, 0},
			&std.Vector3{4, 1.5, 1})
	}
	e.AddComponent(transform)

	quad := render.NewMesh()
	quad.Set(std.QuadMesh)
	e.AddComponent(quad)

	collider := collision.NewCollider()
	collider.OnHit(func(other *engine.Entity) {
		switch other.Tag() {
		case 2: // cable block
			ps := player.Component(&PlayerState{}).(*PlayerState)
			velocity := player.Component(&physics.Velocity{}).(*physics.Velocity)
			// transform := player.Component(&std.Transform{}).(*std.Transform)
			ps.SetRunning(false)
			velocity.SetX(0)
			velocity.SetY(0)
			cableBlocks := ps.CableBlocks()
			for i := 0; i < len(cableBlocks); i++ {
				e.Engine().DeleteEntity(cableBlocks[i])
			}
			ps.ResetLevel()
			// startPos := ps.StartPos()
			// transform.SetPosition(startPos.X, startPos.Y, startPos.Z)
			break
		}
	})
	collider.Set(transform, &std.Rect{&std.Vector2{-1, -1}, 2, 2})
	e.AddComponent(collider)
}

func handleCable(e *engine.Entity) {
	material := render.NewMaterial()
	material.SetColor(
		&std.Color{1.0, 1.0, 1.0, 1.0},
	)
	sa := render.NewImage()
	sa.LoadImage("assets", "cableblock.png")
	s := render.NewSprite(sa)
	s.SetSpriteSize(8, 3)
	material.SetSprite(s)

	ps := e.Component(&PlayerState{}).(*PlayerState)
	cableDrawer := scripting.NewScript()
	ticker := 0.0
	cableDrawer.OnUpdate(func() {
		if ps.Running() {
			if ticker > 60 {
				if math.Mod(ticker, 20) == 0 {
					position := e.Component(&std.Transform{}).(*std.Transform).Position()
					cableBlock := e.Engine().NewEntity()
					cableBlock.AddComponent(material)

					createCableBlock(cableBlock, &std.Vector3{position.X - 3, position.Y - 8, 1}, e, ps.Vertical())
					ps.AddCableBlock(cableBlock)
				}
			}
			ticker++
		}
	})
	e.AddComponent(cableDrawer)
}
