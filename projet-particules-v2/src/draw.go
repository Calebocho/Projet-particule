package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particle"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten. Cette fonction pourra être légèrement modifiée quand
// c'est précisé dans le sujet.
func (g *game) Draw(screen *ebiten.Image) {

	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particle.Particle)
		if ok {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorScale.Scale(float32(p.ColorRed), float32(p.ColorGreen), float32(p.ColorBlue), float32(p.Opacity))
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}

	if config.General.Debug == true {
		if config.General.DynamiqueDuSysDePart == true {
			parametres := fmt.Sprintf(
				"Position de spawn: (%d, %d)\n"+
					"Temps de duree de vie: %.1f\n"+
					"Spawn rate: %.1f\n"+
					"Gravity: %.1f\n",
				config.General.SpawnX,
				config.General.SpawnY,
				config.General.TempsDureeDeVie,
				config.General.SpawnRate,
				config.General.Gravity,
			)
			ebitenutil.DebugPrintAt(screen, parametres, 800, 0)
		}
	}

	if config.General.Debug == true {
		ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
	}

}
