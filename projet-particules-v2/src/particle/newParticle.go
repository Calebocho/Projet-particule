package particle

/**
*Fonction qui renvoie une particule définie par
*
*	PositionX, PositionY            float64 valeur par défaut 0, 0
*	Rotation                        float64 valeur par défaut 0
*	ScaleX, ScaleY                  float64 valeur par défaut 1,1
*	ColorRed, ColorGreen, ColorBlue float64 valeur par défaut 1,1,1
*	Opacity 						float64 valeur par défaut 1
*
*   La particule peut ne pas être à l'écran
*   @param optionals : de 0 à 9 paramètres, les paramètres sont pris à partir de la gauche
*   @return la particule
**/

func NewParticle(optionals ...float64) Particle {
	//définie les valeurs par défaut des paramètres
	params := []float64{0, 0, 0, 1, 1, 1, 1, 1, 1}
	lo := len(optionals)
	for i := 0; i < lo; i++ {
		params[i] = optionals[i]
	}
	p := Particle{
		PositionX:  params[0],
		PositionY:  params[1],
		Rotation:   params[2],
		ScaleX:     params[3],
		ScaleY:     params[4],
		ColorRed:   params[5],
		ColorGreen: params[6],
		ColorBlue:  params[7],
		Opacity:    params[8],
	}

	return p
}
