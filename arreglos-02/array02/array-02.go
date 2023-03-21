/*
2. Escribir una función que reciba dos arreglos (que representan vectores) del mismo tamaño
y devuelvan la suma y el producto escalar de los vectores
*/

package array02

func ProductoYSumaDeVectores(array, array_1 []int) (int, int) {
	suma := 0
	productoEscalar := 0

	for i := 0; i < len(array); i++ {
		for j := i; j < len(array_1); j++ {

			productoEscalar = productoEscalar + array[i]*array_1[j]

			suma = suma + array[i] + array_1[j]
		}
	}
	return productoEscalar, suma
}
