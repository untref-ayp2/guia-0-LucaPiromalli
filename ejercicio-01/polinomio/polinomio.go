package polinomio

import "fmt"

/*1. Definir una función que, dado los coeficientes de un polinomio de grado n (números flotantes),
 muestre por pantalla el polinomio completo, por ejemplo si recibe los coeficientes:
3.0, 2.0 y 1.0 debe mostrar 3.0 + 2.0 X + 1.0 X**2  */

func Polinomio(coeficiente int) (resultado string) {
	var numero int
	for i := 0; i < coeficiente; i++ {

		fmt.Print("Ingrese un entero: ")
		fmt.Scanln(&numero)

		if i == 0 {
			resultado += string(numero) + " "
		} else {
			resultado += string(numero) + "x**" + string(i)
		}
	}
	return
}

/*/ Valores de retorno nombrados
resultado = dividendo / divisor
resto = dividendo % divisor
return*/
