package figuras

import "fmt"

type geometria interface {
	area() float64
	perimetro() float64
}

func Medidas(g geometria) {
	fmt.Println(g)
	fmt.Printf("Área: %9.1f\n", g.area())
	fmt.Printf("Perímetro: %9.1f\n", g.perimetro())
}
