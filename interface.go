package figuras

import "fmt"

type Figura interface {
	calculaArea() float32
	calculaPerimetro() float32
}

func CalculaAreaFigura(figura Figura) {
	result := figura.calculaArea()
	fmt.Println(figura)
	fmt.Println(result)
}

func CalculaPerimetroFigura(figura Figura) {
	result := figura.calculaPerimetro()
	fmt.Println(figura)
	fmt.Println(result)
}
