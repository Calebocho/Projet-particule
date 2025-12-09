package particle

import (
	"math"
	"reflect"
	"testing"
)

func TestNoArgs(t *testing.T) {
	expected := Particle{
		0,
		0,
		0,
		1,
		1,
		1,
		1,
		1,
		1}
	result := NewParticle()
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestOneArg(t *testing.T) {
	expected := Particle{
		2,
		0,
		0,
		1,
		1,
		1,
		1,
		1,
		1}
	result := NewParticle(2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestTwoArgs(t *testing.T) {
	expected := Particle{
		3,
		2,
		0,
		1,
		1,
		1,
		1,
		1,
		1}
	result := NewParticle(3, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestThreeArgs(t *testing.T) {
	expected := Particle{
		3,
		2,
		math.Pi,
		1,
		1,
		1,
		1,
		1,
		1}
	result := NewParticle(3, 2, math.Pi)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestFourArgs(t *testing.T) {
	expected := Particle{
		3,
		2,
		math.Pi,
		2,
		1,
		1,
		1,
		1,
		1}
	result := NewParticle(3, 2, math.Pi, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestFiveArgs(t *testing.T) {
	expected := Particle{
		3,
		2,
		math.Pi / 4,
		0.5,
		2,
		1,
		1,
		1,
		1}
	result := NewParticle(3, 2, math.Pi/4, 0.5, 2)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestSixArgs(t *testing.T) {
	expected := Particle{
		1,
		2,
		0,
		1,
		2,
		0.8,
		1,
		1,
		1,
	}
	result := NewParticle(1, 2, 0, 1, 2, 0.8)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestSevenArgs(t *testing.T) {
	expected := Particle{
		50,
		0,
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		1,
		1}
	result := NewParticle(50, 0, math.Pi, 1.2, 1.1, 0.3, 0.4)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}

func TestEightArgs(t *testing.T) {
	expected := Particle{
		20,
		30,
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		1}
	result := NewParticle(20, 30, math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5)
	if !reflect.DeepEqual(result, expected) {

		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}
func TestTenArgs(t *testing.T) {
	expected := Particle{
		20.0,
		10.0,
		math.Pi,
		1.2,
		1.1,
		0.3,
		0.4,
		0.5,
		2}
	result := NewParticle(20.0, 10.0, math.Pi, 1.2, 1.1, 0.3, 0.4, 0.5, 2)
	if !reflect.DeepEqual(result, expected) {
		t.Fatal("Attendu :", expected, " fourni :", result)
	}
}
