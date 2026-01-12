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

	g.system.Update()

	return nil
}
