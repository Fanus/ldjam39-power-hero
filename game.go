package main

import (
	"github.com/autovelop/playthos"
	"github.com/autovelop/playthos/audio"
	"github.com/autovelop/playthos/audio/wav"
	"github.com/autovelop/playthos/keyboard"
	_ "github.com/autovelop/playthos/opengl"
	_ "github.com/autovelop/playthos/opengl-glfw"
	"github.com/autovelop/playthos/render"
	"github.com/autovelop/playthos/std"
	"os"
)

func main() {
	game := engine.New("LDJAM39", "github.com/autovelop/ldjam39power", &engine.Settings{
		true,
		1024,
		768,
		false,
	})
	if engine.Play() {
		kb := game.Listener(&keyboard.Keyboard{})
		kb.On(keyboard.KeyEscape, func(action ...uint) {
			switch action[0] {
			case keyboard.ActionRelease:
				game.Stop()
				os.Exit(0)
			}
		})

		// running := false
		mu := game.NewEntity()
		musicWAV := wav.NewWAVFile()
		musicWAV.Load("assets", "music3.wav")
		music := audio.NewSound()
		music.Set(musicWAV, true, true)
		mu.AddComponent(music)

		cam := game.NewEntity()
		createCamera(cam)

		levels := []*Level{}

		/*
			LEVELS
		*/
		levels = append(levels, BuildLevel1(game))
		levels = append(levels, BuildLevel2(game))
		levels = append(levels, BuildLevel3(game))
		levels = append(levels, BuildEnd(game))

		/*
			PLAYER
		*/
		player := game.NewEntity()
		createPlayer(player, &std.Vector3{0, 0, 3}, cam, levels)

		handlePlayerInput(kb, player)

		handleCable(player)

		game.Start()
	} else {
		game.Deploy(engine.PlatformLinux, engine.PlatformMacOS, engine.PlatformWindows)
		// game.Deploy(engine.PlatformLinux, engine.PlatformWindows)
		// game.Deploy(engine.PlatformLinux)
	}
}

func BuildLevel1(game *engine.Engine) *Level {
	offset := float32(0)
	level := &Level{}
	level.SetCameraPosition(&std.Vector3{offset, 0, 0})

	terrain := game.NewEntity()
	createTerrain(terrain, &std.Vector3{offset, 0, -1})

	startZone := game.NewEntity()
	createStartZone(startZone, &std.Vector3{offset - 116, 0, 2})

	safeZone := game.NewEntity()
	createSafeZone(safeZone, &std.Vector3{offset + 116, 0, 2})

	nextN := game.NewEntity()
	createNextZone(nextN, &std.Vector3{offset + 116, 192, 3})

	nextS := game.NewEntity()
	createNextZone(nextS, &std.Vector3{offset + 116, -192, 3})

	enemy1 := game.NewEntity()
	w1 := []*Waypoint{}
	w1 = append(w1, &Waypoint{&std.Vector3{-0.2, 0, 0}, &std.Vector3{-60, 0, 0}})
	w1 = append(w1, &Waypoint{&std.Vector3{0, -0.2, 0}, &std.Vector3{0, -80, 0}})
	w1 = append(w1, &Waypoint{&std.Vector3{0, 0.2, 0}, &std.Vector3{0, 60, 0}})
	w1 = append(w1, &Waypoint{&std.Vector3{0.2, 0, 0}, &std.Vector3{60, 0, 0}})
	createEnemy(enemy1, &std.Vector3{60, 60, 0}, &std.Vector2{-300, 300}, w1)

	enemy2 := game.NewEntity()
	w2 := []*Waypoint{}
	w2 = append(w2, &Waypoint{&std.Vector3{0, -0.2, 0}, &std.Vector3{0, 0, 0}})
	w2 = append(w2, &Waypoint{&std.Vector3{0, 0.2, 0}, &std.Vector3{0, 90, 0}})
	createEnemy(enemy2, &std.Vector3{90, 90, 0}, &std.Vector2{-300, 300}, w2)

	enemy3 := game.NewEntity()
	w3 := []*Waypoint{}
	w3 = append(w3, &Waypoint{&std.Vector3{0, -0.2, 0}, &std.Vector3{0, -90, 0}})
	w3 = append(w3, &Waypoint{&std.Vector3{0, 0.2, 0}, &std.Vector3{0, 16, 0}})
	createEnemy(enemy3, &std.Vector3{40, 13, 0}, &std.Vector2{-300, 300}, w3)

	level.AddEnemies(enemy1, enemy2, enemy3)
	return level
}

func BuildLevel2(game *engine.Engine) *Level {
	offset := float32(300)
	level := &Level{}
	level.SetCameraPosition(&std.Vector3{offset, 0, 0})

	terrain := game.NewEntity()
	createTerrain(terrain, &std.Vector3{offset, 0, -1})

	startZone := game.NewEntity()
	createStartZone(startZone, &std.Vector3{offset - 116, 0, 2})

	safeZone := game.NewEntity()
	createSafeZone(safeZone, &std.Vector3{offset + 116, 0, 2})

	nextN := game.NewEntity()
	createNextZone(nextN, &std.Vector3{offset + 116, 192, 3})

	nextS := game.NewEntity()
	createNextZone(nextS, &std.Vector3{offset + 116, -192, 3})

	enemy1 := game.NewEntity()
	w1 := []*Waypoint{}
	w1 = append(w1, &Waypoint{&std.Vector3{0, -0.2, 0}, &std.Vector3{0, -70, 0}})
	w1 = append(w1, &Waypoint{&std.Vector3{0.2, 0, 0}, &std.Vector3{offset + 80, 0, 0}})
	w1 = append(w1, &Waypoint{&std.Vector3{0, -0.2, 0}, &std.Vector3{0, -90, 0}})
	w1 = append(w1, &Waypoint{&std.Vector3{0, 0.2, 0}, &std.Vector3{0, -70, 0}})
	w1 = append(w1, &Waypoint{&std.Vector3{-0.2, 0, 0}, &std.Vector3{offset - 80, 0, 0}})
	w1 = append(w1, &Waypoint{&std.Vector3{0, 0.2, 0}, &std.Vector3{0, 80, 0}})
	createEnemy(enemy1, &std.Vector3{offset - 80, 80, 0}, &std.Vector2{-300, 300}, w1)

	enemy2 := game.NewEntity()
	w2 := []*Waypoint{}
	w2 = append(w2, &Waypoint{&std.Vector3{0, 0.2, 0}, &std.Vector3{0, 90, 0}})
	w2 = append(w2, &Waypoint{&std.Vector3{0, -0.2, 0}, &std.Vector3{0, 0, 0}})
	createEnemy(enemy2, &std.Vector3{offset, 0, 0}, &std.Vector2{-300, 300}, w2)

	enemy3 := game.NewEntity()
	w3 := []*Waypoint{}
	w3 = append(w3, &Waypoint{&std.Vector3{0, 0.2, 0}, &std.Vector3{0, 70, 0}})
	w3 = append(w3, &Waypoint{&std.Vector3{0, -0.2, 0}, &std.Vector3{0, -50, 0}})
	createEnemy(enemy3, &std.Vector3{offset + 50, -60, 0}, &std.Vector2{-300, 300}, w3)

	enemy4 := game.NewEntity()
	w4 := []*Waypoint{}
	w4 = append(w4, &Waypoint{&std.Vector3{0, 0.2, 0}, &std.Vector3{0, 40, 0}})
	w4 = append(w4, &Waypoint{&std.Vector3{0, -0.2, 0}, &std.Vector3{0, -90, 0}})
	createEnemy(enemy4, &std.Vector3{offset + 70, -90, 0}, &std.Vector2{-300, 300}, w4)

	level.AddEnemies(enemy1, enemy2, enemy3, enemy4)
	return level
}

func BuildLevel3(game *engine.Engine) *Level {
	offset := float32(600)
	level := &Level{}
	level.SetCameraPosition(&std.Vector3{offset, 0, 0})

	terrain := game.NewEntity()
	createTerrain(terrain, &std.Vector3{offset, 0, -1})

	startZone := game.NewEntity()
	createStartZone(startZone, &std.Vector3{offset - 116, 0, 2})

	safeZone := game.NewEntity()
	createSafeZone(safeZone, &std.Vector3{offset + 116, 0, 2})

	nextN := game.NewEntity()
	createNextZone(nextN, &std.Vector3{offset + 116, 192, 3})

	nextS := game.NewEntity()
	createNextZone(nextS, &std.Vector3{offset + 116, -192, 3})

	enemy1 := game.NewEntity()
	w1 := []*Waypoint{}
	w1 = append(w1, &Waypoint{&std.Vector3{0, 0.3, 0}, &std.Vector3{0, 70, 0}})
	w1 = append(w1, &Waypoint{&std.Vector3{0, -0.3, 0}, &std.Vector3{0, -90, 0}})
	createEnemy(enemy1, &std.Vector3{offset - 70, -90, 0}, &std.Vector2{-300, 300}, w1)

	enemy2 := game.NewEntity()
	w2 := []*Waypoint{}
	w2 = append(w2, &Waypoint{&std.Vector3{0, -0.3, 0}, &std.Vector3{0, -70, 0}})
	w2 = append(w2, &Waypoint{&std.Vector3{0, 0.3, 0}, &std.Vector3{0, 90, 0}})
	createEnemy(enemy2, &std.Vector3{offset - 40, 90, 0}, &std.Vector2{-300, 300}, w2)

	enemy3 := game.NewEntity()
	w3 := []*Waypoint{}
	w3 = append(w3, &Waypoint{&std.Vector3{0, 0.3, 0}, &std.Vector3{0, 70, 0}})
	w3 = append(w3, &Waypoint{&std.Vector3{0, -0.3, 0}, &std.Vector3{0, -90, 0}})
	createEnemy(enemy3, &std.Vector3{offset - 10, -90, 0}, &std.Vector2{-300, 300}, w3)

	enemy4 := game.NewEntity()
	w4 := []*Waypoint{}
	w4 = append(w4, &Waypoint{&std.Vector3{0, -0.3, 0}, &std.Vector3{0, -70, 0}})
	w4 = append(w4, &Waypoint{&std.Vector3{0, 0.3, 0}, &std.Vector3{0, 90, 0}})
	createEnemy(enemy4, &std.Vector3{offset + 20, 90, 0}, &std.Vector2{-300, 300}, w4)

	enemy5 := game.NewEntity()
	w5 := []*Waypoint{}
	w5 = append(w5, &Waypoint{&std.Vector3{0, 0.3, 0}, &std.Vector3{0, 70, 0}})
	w5 = append(w5, &Waypoint{&std.Vector3{0, -0.3, 0}, &std.Vector3{0, -90, 0}})
	createEnemy(enemy5, &std.Vector3{offset + 50, -90, 0}, &std.Vector2{-300, 300}, w5)

	enemy6 := game.NewEntity()
	w6 := []*Waypoint{}
	w6 = append(w6, &Waypoint{&std.Vector3{0, -0.3, 0}, &std.Vector3{0, -70, 0}})
	w6 = append(w6, &Waypoint{&std.Vector3{0, 0.3, 0}, &std.Vector3{0, 90, 0}})
	createEnemy(enemy6, &std.Vector3{offset + 80, 90, 0}, &std.Vector2{-300, 300}, w6)

	level.AddEnemies(enemy1)
	return level
}

func BuildEnd(game *engine.Engine) *Level {
	offset := float32(900)
	level := &Level{}
	level.SetCameraPosition(&std.Vector3{offset, 0, 0})

	end := game.NewEntity()
	createEnd(end, &std.Vector3{offset, 0, -1})
	return level
}

func createEnd(e *engine.Entity, p *std.Vector3) {
	transform := std.NewTransform()
	transform.Set(
		p,
		&std.Vector3{0, 0, 0},
		&std.Vector3{256, 192, 1})
	e.AddComponent(transform)

	quad := render.NewMesh()
	quad.Set(std.QuadMesh)
	e.AddComponent(quad)

	material := render.NewMaterial()
	material.SetColor(
		&std.Color{1.0, 1.0, 1.0, 1.0},
	)

	sa := render.NewImage()
	sa.LoadImage("assets", "Outro.png")
	s := render.NewSprite(sa)
	s.SetSpriteSize(256, 192)
	material.SetSprite(s)

	e.AddComponent(material)
}
