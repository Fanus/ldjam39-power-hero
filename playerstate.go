package main

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/animation"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
)

type PlayerState struct {
	engine.Component
	cableBlocks []*engine.Entity
	startPos    *std.Vector3

	camera *engine.Entity
	entity *engine.Entity

	walkingWSprite    *render.Texture
	walkingDownSprite *render.Texture
	walkingUpSprite   *render.Texture
	idleSprite        *render.Texture

	idleAnimation        *animation.AnimationClip
	walkingWAnimation    *animation.AnimationClip
	walkingUpAnimation   *animation.AnimationClip
	walkingDownAnimation *animation.AnimationClip

	running  bool
	vertical bool

	levels       []*Level
	currentLevel uint
}

func (p *PlayerState) Set(cb []*engine.Entity, s *std.Vector3, r bool, e *engine.Entity, c *engine.Entity, ls []*Level) {
	p.cableBlocks = cb
	p.startPos = s
	p.running = r
	p.entity = e
	p.camera = c
	p.levels = ls
	p.currentLevel = 0
	if len(p.levels) > 0 {
		p.levels[p.currentLevel].Load(p)
	}
}

func (p *PlayerState) Camera() *engine.Entity {
	return p.camera
}

func (p *PlayerState) Player() *engine.Entity {
	return p.entity
}

func (p *PlayerState) Running() bool {
	return p.running
}

func (p *PlayerState) SetRunning(r bool) {
	p.running = r
}

func (p *PlayerState) Vertical() bool {
	return p.vertical
}

func (p *PlayerState) SetVertical(v bool) {
	p.vertical = v
}

func (p *PlayerState) AddCableBlock(e *engine.Entity) {
	p.cableBlocks = append(p.cableBlocks, e)
}

func (p *PlayerState) CableBlocks() []*engine.Entity {
	return p.cableBlocks
}

func (p *PlayerState) StartPos() *std.Vector3 {
	return p.startPos
}

func (p *PlayerState) SetStartPos(pos *std.Vector3) {
	p.startPos = pos
}

func (p *PlayerState) Idle() (*render.Texture, *animation.AnimationClip) {
	return p.idleSprite, p.idleAnimation
}

func (p *PlayerState) WalkingW() (*render.Texture, *animation.AnimationClip) {
	return p.walkingWSprite, p.walkingWAnimation
}

func (p *PlayerState) WalkingUp() (*render.Texture, *animation.AnimationClip) {
	return p.walkingUpSprite, p.walkingUpAnimation
}

func (p *PlayerState) WalkingDown() (*render.Texture, *animation.AnimationClip) {
	return p.walkingDownSprite, p.walkingDownAnimation
}

func (p *PlayerState) Completed() bool {
	return p.levels[p.currentLevel].Completed()
}

func (p *PlayerState) SetLevelCompleted(l uint) {
	p.levels[l].SetCompleted(true)
}

func (p *PlayerState) CurrentLevel() uint {
	return p.currentLevel
}

func (p *PlayerState) ResetLevel() {
	p.levels[p.currentLevel].Load(p)
}

func (p *PlayerState) NextLevel() {
	p.currentLevel++
	if int(p.currentLevel) <= len(p.levels)-1 {
		p.levels[p.currentLevel].Load(p)
	}
}

func (p *PlayerState) SetSprites(i *render.Texture, w *render.Texture, u *render.Texture, d *render.Texture) {
	p.idleSprite = i
	p.walkingWSprite = w
	p.walkingUpSprite = u
	p.walkingDownSprite = d
}

func (p *PlayerState) SetAnimations(i *animation.AnimationClip, w *animation.AnimationClip, u *animation.AnimationClip, d *animation.AnimationClip) {
	p.idleAnimation = i
	p.walkingWAnimation = w
	p.walkingUpAnimation = u
	p.walkingDownAnimation = d
}
