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
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*particle.Particle)
		p.SpeedY += config.General.Gravity
		p.PositionX = p.PositionX + p.SpeedX
		p.PositionY = p.PositionY + p.SpeedY

	}
}
