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

	// extension 5.6
	// ici je convertis la vitesse en entier pour pouvoir l'additionner aux coordonnées
	speed := int(config.General.MoveSpeedGenerator)

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		// ici je déplace le générateur vers la gauche
		config.General.SpawnX -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		// ici je déplace le générateur vers la droite
		config.General.SpawnX += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		// ici je déplace le générateur vers le haut
		config.General.SpawnY -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		// ici je déplace le générateur vers le bas
		config.General.SpawnY += speed
	}
	// extension 5.7
	if config.General.DynamiqueDuSysDePart == true {
		// ici je peux augmenter ou diminuer la duree de vie
		if ebiten.IsKeyPressed(ebiten.KeyEqual) == true {
			config.General.TempsDureeDeVie = config.General.TempsDureeDeVie + 5
		}
		if ebiten.IsKeyPressed(ebiten.KeyMinus) && config.General.TempsDureeDeVie > 0 {
			config.General.TempsDureeDeVie = config.General.TempsDureeDeVie - 5
		}
		// Je m'assure qu'elle ne reste jamais en dessous de 0
		if config.General.TempsDureeDeVie < 0 {
			config.General.TempsDureeDeVie = 0
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
		//ici je peux augmenter ou diminuer le rayon de cercle
		if ebiten.IsKeyPressed(ebiten.KeyA) == true {
			config.General.Generateurtaille = config.General.Generateurtaille + 5
		}
		if ebiten.IsKeyPressed(ebiten.KeyZ) == true {
			config.General.Generateurtaille = config.General.Generateurtaille - 5
		}

	}

	g.system.Update()

	return nil
}
