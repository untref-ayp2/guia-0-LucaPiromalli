package polinomio

import "fmt"

/*1. Definir una función que, dado los coeficientes de un polinomio de grado n (números flotantes),
 muestre por pantalla el polinomio completo, por ejemplo si recibe los coeficientes:
3.0, 2.0 y 1.0 debe mostrar 3.0 + 2.0 X + 1.0 X**2  */

func Polinomio(coeficiente ...float32) (resultado string) {
	arreglo := make([]string, len(coeficiente)) // array del tamaño de la cantidad de coeficientes
	for i := 0; i < len(coeficiente); i++ {

		if i > 1 {
			arreglo[i] = fmt.Sprintf(" %+.1f X**%v ", coeficiente[i], i) // (%+ .1f) Con este ejemplo el input sería "-3.1416" y el output sería "- 3.1 ""
		} else if i == 1 {
			arreglo[i] = fmt.Sprintf(" %+.1f X ", coeficiente[i])
		} else {
			arreglo[i] = fmt.Sprintf(" %.1f ", coeficiente[i])
		}
	}
	for _, coeficiente := range arreglo {
		fmt.Print(coeficiente)
	}
}

/*/ Valores de retorno nombrados
resultado = dividendo / divisor
resto = dividendo % divisor
return*/
