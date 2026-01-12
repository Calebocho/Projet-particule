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

//TEST 3 : Mouvement
// Vérifie que la fonction Update déplace bien les particules selon leur vitesse

func TestUpdate_Movement(t *testing.T) {
	resetConfig()
	config.General.InitNumParticles = 1
	config.General.SpawnRate = 0

	s := NewSystem()
	p := s.Content.Front().Value.(*particle.Particle)

	// On force une position et une vitesse connues pour le test
	p.PositionX = 100
	p.PositionY = 100
	p.SpeedY = 5
	p.SpeedX = 10

	// On lance un tour de mise à jour
	s.Update()

	// Nouvelle position attendue : 100 + 10 = 110, 100 + 5 = 105
	if p.PositionX != 110 || p.PositionY != 105 {
		t.Errorf("Mouvement incorrect : attendu (110, 105), obtenu (%f, %f)", p.PositionX, p.PositionY)
	}
}

// TEST 4 : Génération Continue
// Vérifie que le système génère des particules petit à petit
func TestUpdate_SpawnRate(t *testing.T) {
	resetConfig()
	config.General.InitNumParticles = 0
	config.General.SpawnRate = 0.5 // 0.5 particule par tour

	s := NewSystem()

	// Tour 1 : Tirelire = 0.5 -> Pas de création
	s.Update()
	if s.Content.Len() != 0 {
		t.Errorf("Tour 1 : ne devrait pas créer de particule avec SpawnRate 0.5. Particules : %d", s.Content.Len())
	}

	// Tour 2 : Tirelire = 1.0 -> Création d'une particule !
	s.Update()
	if s.Content.Len() != 1 {
		t.Errorf("Tour 2 : devrait avoir créé 1 particule. Particules : %d", s.Content.Len())
	}

	// Tour 3 : Tirelire = 0.5 (après avoir payé) -> Pas de création
	s.Update()
	if s.Content.Len() != 1 {
		t.Errorf("Tour 3 : ne devrait pas créer de nouvelle particule. Particules : %d", s.Content.Len())
	}
}

// TEST 5 : Extension Gravité
/* Vérifie que la gravité augmente bien la vitesse Y de la particule avec une Gravité forte pas de nouvelles particules et on empêche la mort par opacité
on réxupère une particule on sauvegarde sa vitesse initiale avec VitesseInit et on fait un tour de boucle la vitesse attendue est VitesseInit + Gravité
*/
func TestGravity(t *testing.T) {
	resetConfig()
	config.General.InitNumParticles = 1
	config.General.Gravity = 10.0
	config.General.SpawnRate = 0
	config.General.OpacitySpeed = 0

	s := NewSystem()
	p := s.Content.Front().Value.(*particle.Particle)

	vitesseInit := p.SpeedY

	s.Update()

	vitesseAttendue := vitesseInit + config.General.Gravity

	if p.SpeedY != vitesseAttendue {
		t.Errorf("La gravité ne s'applique pas. VitesseY avant: %f, après: %f (Attendu: %f)", vitesseInit, p.SpeedY, vitesseAttendue)
	}
}

// TEST 6 : Extension Sortie d'Écran
/* Vérifie qu'une particule est supprimée si elle sort de la fenêtre en plaçant hors de l'écran une particule et en vérifiant si elle est bien suprimée de la liste après un update pour celà on définie une petite fenêtre et on tp la particule très loin avant l'update elle est présente et après elle disparait

 */
func TestScreen(t *testing.T) {
	resetConfig()
	config.General.InitNumParticles = 1
	config.General.SpawnRate = 0

	config.General.WindowSizeX = 100
	config.General.WindowSizeY = 100

	s := NewSystem()
	p := s.Content.Front().Value.(*particle.Particle)

	p.PositionX = -500

	if s.Content.Len() != 1 {
		t.Error("La particule devrait être là avant l'update")
	}
	s.Update()

	if s.Content.Len() != 0 {
		t.Errorf("La particule hors écran n'a pas été supprimée. Reste : %d", s.Content.Len())
	}
}

// TEST 7 : Extension Opacité et Durée de Vie
/* Vérifie que l'opacité baisse et que la particule meurt à 0 ( quand elle est invisible) On configure une baisse d'opacité drastique pour tuer la particule en 1 coup
 */
func TestOpacityDeath(t *testing.T) {
	resetConfig()
	config.General.InitNumParticles = 1
	config.General.SpawnRate = 0

	config.General.OpacitySpeed = 1.1

	s := NewSystem()

	p := s.Content.Front().Value.(*particle.Particle)
	if p.Opacity != 1 {
		t.Error("L'opacité initiale devrait être à 1")
	}
	s.Update()

	if s.Content.Len() != 0 {
		t.Errorf("La particule invisible n'a pas été supprimée. Reste : %d", s.Content.Len())
	}
}
