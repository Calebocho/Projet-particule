package main

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction ne devrait pas être modifiée sauf
// pour les deux dernières extensions.
func (g *game) Update() error {

	speed := int(config.General.MoveSpeedGenerator)
	// extension 5.6
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		config.General.SpawnX -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		config.General.SpawnX += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		config.General.SpawnY -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		config.General.SpawnY += speed
	}
	// extension 5.7
	if config.General.DynamiqueDuSysDePart == true {
		// ici je peux augmenter ou diminuer la duree de vie
		if ebiten.IsKeyPressed(ebiten.KeyEqual) == true {
			config.General.TempsDureeDeVie = config.General.TempsDureeDeVie + 5
		}
		if config.General.TempsDureeDeVie < 0 {
			config.General.TempsDureeDeVie = 0
		} else if ebiten.IsKeyPressed(ebiten.KeyMinus) == true {
			config.General.TempsDureeDeVie = config.General.TempsDureeDeVie - 5
		}
		// ici je prend la position de spawn avec un cursor
		X, Y := ebiten.CursorPosition()
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			config.General.SpawnX = X
			config.General.SpawnY = Y
		}
		//ici je peux augmenter ou diminuer la gravite
		if ebiten.IsKeyPressed(ebiten.KeyW) == true {
			config.General.Gravity = config.General.Gravity + 0.1
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) == true {
			config.General.Gravity = config.General.Gravity - 0.1
		}

	}

	g.system.Update()

	return nil
}
