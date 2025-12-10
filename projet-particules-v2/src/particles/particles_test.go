package particles

import (
	"project-particles/config"
	"project-particles/particle"
	"testing"
)

//ResetConfig qui remet la configuration à zéro pour chaque test pour éviter tout mélange entre test

func resetConfig() {
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 800
	config.General.RandomSpawn = false
	config.General.InitNumParticles = 0
	config.General.Gravity = 0
	config.General.SpawnX = 400
	config.General.SpawnY = 400
}

//TEST 1
// Vérifie que NewSystem crée bien le nombre exact de particules demandé

func TestNewSystem(t *testing.T) {
	resetConfig()
	config.General.InitNumParticles = 100

	s := NewSystem()

	if s.Content.Len() != 100 {
		t.Errorf("Initialisation incorrecte : on attent 100 particules,mais là il y'a %d", s.Content.Len())
	}
}

//TEST 2 : Position fixe
// Vérifie que RandomSpawn est faux , les particules naissent bien aux coordonnées SpawnX/SpawnY

func TestNewSystem_PosFix(t *testing.T) {
	resetConfig()
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = false
	config.General.SpawnX = 250
	config.General.SpawnY = 450

	s := NewSystem()

	p := s.Content.Front().Value.(*particle.Particle)

	if p.PositionY != 450 || p.PositionX != 250 {
		t.Errorf("Position fixe incorrecte : attendu (250, 450), obtenu (%f, %f)", p.PositionX, p.PositionY)
	}
}
