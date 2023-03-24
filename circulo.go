package figuras

import "math"

type Circulo struct {
	Radio float32
}

func (circulo *Circulo) calculaArea() float32 {
	return math.Pi * circulo.Radio * circulo.Radio
}

func (circulo *Circulo) calculaPerimetro() float32 {
	return math.Pi * circulo.Radio * 2
}
