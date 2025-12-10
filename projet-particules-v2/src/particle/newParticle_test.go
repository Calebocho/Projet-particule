package particle

import (
	"math"
	"project-particles/config"
	"reflect"
	"testing"
)

// Configure les paramÃ¨tres de base pour les tests (taille de l'Ã©cran)
func configurerTest() {
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
}

// --- Tests sur le nombre d'arguments ---

// Test de la crÃ©ation d'une particule sans argument (valeurs par dÃ©faut).
func TestNoArgs(t *testing.T) {
	configurerTest()
	config.General.RandomSpawn = false
	expected := Particle{
		PositionX:  float64(config.General.WindowSizeX),
		PositionY:  float64(config.General.WindowSizeY),
		Rotation:   0,
		ScaleX:     1,
		ScaleY:     1,
		ColorRed:   1,
		ColorGreen: 1,
		ColorBlue:  1,
		Opacity:    1,
		SpeedX:     1,
		SpeedY:     1,
	}
	result := NewParticle()
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la création d'une particule avec 1 argument (Ã©crase Opacity).
func TestOneArg(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  700,
		PositionY:  float64(config.General.WindowSizeY),
		Rotation:   0,
		ScaleX:     1,
		ScaleY:     1,
		ColorRed:   1,
		ColorGreen: 1,
		ColorBlue:  1,
		Opacity:    1,
		SpeedX:     1,
		SpeedY:     1,
	}
	result := NewParticle(700)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la crÃ©ation d'une particule avec 2 arguments (Ã©crase ColorBlue et Opacity).
func TestTwoArgs(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  600,
		PositionY:  500,
		Rotation:   0,
		ScaleX:     1,
		ScaleY:     1,
		ColorRed:   1,
		ColorGreen: 1,
		ColorBlue:  1,
		Opacity:    1,
		SpeedX:     1,
		SpeedY:     1,
	}
	result := NewParticle(600, 500)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la crÃ©ation d'une particule avec 3 arguments.
func TestThreeArgs(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  600,
		PositionY:  500,
		Rotation:   math.Pi,
		ScaleX:     1,
		ScaleY:     1,
		ColorRed:   1,
		ColorGreen: 1,
		ColorBlue:  1,
		Opacity:    1,
		SpeedX:     1,
		SpeedY:     1,
	}
	result := NewParticle(600, 500, math.Pi)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la crÃ©ation d'une particule avec 4 arguments.
func TestFourArgs(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  600,
		PositionY:  500,
		Rotation:   math.Pi,
		ScaleX:     3,
		ScaleY:     1,
		ColorRed:   1,
		ColorGreen: 1,
		ColorBlue:  1,
		Opacity:    1,
		SpeedX:     1,
		SpeedY:     1,
	}
	result := NewParticle(600, 500, math.Pi, 3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la crÃ©ation d'une particule avec 5 arguments.
func TestFiveArgs(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  600,
		PositionY:  500,
		Rotation:   math.Pi,
		ScaleX:     3,
		ScaleY:     2,
		ColorRed:   1,
		ColorGreen: 1,
		ColorBlue:  1,
		Opacity:    1,
		SpeedX:     1,
		SpeedY:     1,
	}
	result := NewParticle(600, 500, math.Pi, 3, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la crÃ©ation d'une particule avec 6 arguments.
func TestSixArgs(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  600,
		PositionY:  500,
		Rotation:   math.Pi,
		ScaleX:     3,
		ScaleY:     2,
		ColorRed:   0.4,
		ColorGreen: 1,
		ColorBlue:  1,
		Opacity:    1,
		SpeedX:     1,
		SpeedY:     1,
	}
	result := NewParticle(600, 500, math.Pi, 3, 2, 0.4)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la crÃ©ation d'une particule avec 7 arguments.
func TestSevenArgs(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  600,
		PositionY:  500,
		Rotation:   math.Pi,
		ScaleX:     3,
		ScaleY:     2,
		ColorRed:   0.4,
		ColorGreen: 0.3,
		ColorBlue:  1,
		Opacity:    1,
		SpeedX:     1,
		SpeedY:     1,
	}
	result := NewParticle(600, 500, math.Pi, 3, 2, 0.4, 0.3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la crÃ©ation d'une particule avec 8 arguments.
func TestEightArgs(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  600,
		PositionY:  500,
		Rotation:   math.Pi,
		ScaleX:     3,
		ScaleY:     2,
		ColorRed:   0.4,
		ColorGreen: 0.3,
		ColorBlue:  0.8,
		Opacity:    1,
		SpeedX:     1,
		SpeedY:     1,
	}
	result := NewParticle(600, 500, math.Pi, 3, 2, 0.4, 0.3, 0.8)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la crÃ©ation d'une particule avec 9 arguments .
func TestNineArgs(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  600,
		PositionY:  500,
		Rotation:   math.Pi,
		ScaleX:     3,
		ScaleY:     2,
		ColorRed:   0.4,
		ColorGreen: 0.3,
		ColorBlue:  0.8,
		Opacity:    0.7,
		SpeedX:     1,
		SpeedY:     1,
	}
	result := NewParticle(600, 500, math.Pi, 3, 2, 0.4, 0.3, 0.8, 0.7)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la crÃ©ation d'une particule avec 10 arguments .
func TestTenArgs(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  600,
		PositionY:  500,
		Rotation:   math.Pi,
		ScaleX:     3,
		ScaleY:     2,
		ColorRed:   0.4,
		ColorGreen: 0.3,
		ColorBlue:  0.8,
		Opacity:    0.7,
		SpeedX:     1.3,
		SpeedY:     1,
	}
	result := NewParticle(600, 500, math.Pi, 3, 2, 0.4, 0.3, 0.8, 0.7, 1.3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}

// Test de la crÃ©ation d'une particule avec 10 arguments .
func TestElevenArgs(t *testing.T) {
	configurerTest()

	expected := Particle{
		PositionX:  600,
		PositionY:  500,
		Rotation:   math.Pi,
		ScaleX:     3,
		ScaleY:     2,
		ColorRed:   0.4,
		ColorGreen: 0.3,
		ColorBlue:  0.8,
		Opacity:    0.7,
		SpeedX:     1.3,
		SpeedY:     2,
	}
	result := NewParticle(600, 500, math.Pi, 3, 2, 0.4, 0.3, 0.8, 0.7, 1.3, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Attendu: %+v, Obtenu: %+v", expected, result)
	}
}
