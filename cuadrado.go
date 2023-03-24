package figuras

type Cuadrado struct {
	Ancho float32
	Largo float32
}

func (cuadrado *Cuadrado) calculaPerimetro() float32 {
	return cuadrado.Ancho*2 + cuadrado.Largo*2
}

func (cuadrado *Cuadrado) calculaArea() float32 {
	return cuadrado.Ancho * cuadrado.Largo
}

func HolaMundo() {
	fmt.Println("Hola desde modulo Github")
}
