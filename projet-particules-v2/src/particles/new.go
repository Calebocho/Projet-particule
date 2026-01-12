package particles

import (
	"container/list"
	"math"
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
	s := System{Content: list.New()}

	for i := 0; i < config.General.InitNumParticles; i++ {
		s.CreateParticle()
	}

	return s
}
func (s *System) CreateParticle() {
	p := particle.NewParticle(10, 10)
	//extension 5.5
	if config.General.Generateur == "cercle" {
		angle := (rand.Float64() - 0.5) * 2 * math.Pi
		// une formule de math de diff de coordonnes sur oY et oX   pour un cercle
		p.PositionX = float64(config.General.SpawnX) + config.General.Generateurtaille*math.Cos(angle)
		p.PositionY = float64(config.General.SpawnY) + config.General.Generateurtaille*math.Sin(angle)
		// des vitesses initiales des particules en fonction de la forme du générateur
		p.SpeedX = config.General.VitesseInit * math.Cos(angle)
		p.SpeedY = config.General.VitesseInit * math.Sin(angle)

	} else {
		// une position
		if config.General.RandomSpawn {
			p.PositionX = rand.Float64() * float64(config.General.WindowSizeX)
			p.PositionY = rand.Float64() * float64(config.General.WindowSizeY)
		} else {
			p.PositionX = float64(config.General.SpawnX)
			p.PositionY = float64(config.General.SpawnY)
		}

		p.SpeedX = rand.Float64()*10 - 5
		p.SpeedY = rand.Float64()*10 - 5
	}

	if config.General.RandColor {
		p.ColorRed = rand.Float64()
		p.ColorBlue = rand.Float64()
		p.ColorGreen = rand.Float64()
	} else {
		p.ColorRed = config.General.Red
		p.ColorBlue = config.General.Blue
		p.ColorGreen = config.General.Green

	}
	p.ScaleX = 1
	p.ScaleY = 1

	p.Opacity = 1

	//extension 5.3
	if config.General.DureeDeVie == true {
		p.TempsDureeDeVie = config.General.TempsDureeDeVie + float64(rand.Intn(50))
		// J'ai ajouté une fonction aléatoire pour que les particules disparaissent à des moments différents et non en même temps.
	} else {
		p.TempsDureeDeVie = 9999
		// Si nous  n'avons pas de durée de vie, cela signifie que les particules ne disparaîtront pas. et 9999 comme l'infinie
	}

	s.Content.PushFront(&p)
}
