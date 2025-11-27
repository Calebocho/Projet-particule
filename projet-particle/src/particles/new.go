package particles

import (
	"container/list"
	"math/rand"
	"project-particles/config"
	"project-particles/particle"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	// On initialise le système d'abord avec une liste vide
	s := System{Content: list.New()}

	for i := 0; i < config.General.InitNumParticles; i++ {
		s.CreateParticle()
	}

	return s
}
func (s *System) CreateParticle() {
	p := particle.NewParticle()

	// 1. Position
	if config.General.RandomSpawn {
		p.PositionX = rand.Float64() * float64(config.General.WindowSizeX)
		p.PositionY = rand.Float64() * float64(config.General.WindowSizeY)
	} else {
		p.PositionX = float64(config.General.SpawnX)
		p.PositionY = float64(config.General.SpawnY)
	}

	// 2. Vitesse
	p.SpeedX = rand.Float64()*10 - 5
	p.SpeedY = rand.Float64()*10 - 5

	// 3. Apparence
	p.ScaleX, p.ScaleY = 1, 1
	p.ColorRed, p.ColorGreen, p.ColorBlue = 1, 1, 1
	p.Opacity = 1

	// 4. Ajout à la liste
	s.Content.PushFront(&p)
}
