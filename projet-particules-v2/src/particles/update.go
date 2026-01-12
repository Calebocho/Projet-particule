package particles

import (
	"project-particles/config"
	"project-particles/particle"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	s.SpawnRemainder += config.General.SpawnRate

	for s.SpawnRemainder >= 1 {
		s.CreateParticle()
		s.SpawnRemainder -= 1
	}
	e := s.Content.Front()
	for e != nil {
		suivant := e.Next()
		p := e.Value.(*particle.Particle)
		//extension 5.1
		p.SpeedY = p.SpeedY + config.General.Gravity
		p.PositionX = p.PositionX + p.SpeedX
		p.PositionY = p.PositionY + p.SpeedY
		//extension 5.3
		if config.General.DureeDeVie == true {
			p.TempsDureeDeVie = p.TempsDureeDeVie - 1

		}
		//Extension 5.4
		//Ici je rends la particule un peu plus transparente
		p.Opacity -= config.General.OpacitySpeed

		//Extension 5.8
		// Ici je supprime la particule dans le cas ou elle est invisible ou si le temps est dépassé
		if p.Opacity <= 0 || (config.General.DureeDeVie && p.TempsDureeDeVie <= 0) {
			s.Content.Remove(e)
			e = suivant
			continue
		}

		//Extension 5.2
		margin := 10.0
		maxX := float64(config.General.WindowSizeX) + margin
		maxY := float64(config.General.WindowSizeY) + margin

		// ici je vérifie si la particule est sortie du cadre pour la supprimer
		if p.PositionX < -margin || p.PositionX > maxX || p.PositionY < -margin || p.PositionY > maxY {
			s.Content.Remove(e)

		}
		e = suivant
	}
}
