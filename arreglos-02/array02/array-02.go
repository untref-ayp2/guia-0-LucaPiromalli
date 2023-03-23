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

/*
Escribir una función que reciba dos arreglos A y B, de N y M elementos respectivamente
que representan conjuntos de enteros y devuelva una arreglo con la unión y otro con la intersección de A y B.
​*/

func ejercicioTresArray(arrayA []int, arrayB []int) ([]int, []int) {

	union := Union(arrayA, arrayB)

	interseccion := Interseccion(arrayA, arrayB)

	return union, interseccion
}

func Union(arrayA []int, arrayB []int) []int {

	unionDeArrays := arrayA[0:]

	unionDeArrays = append(unionDeArrays, arrayB...)

	return unionDeArrays
}

func Interseccion(arrayA []int, arrayB []int) []int {

	interseccionDeArrays := arrayA[0:]

	for _, numeroA := range arrayA {
		for _, numeroB := range arrayB {

			if numeroA == numeroB {
				interseccionDeArrays = append(interseccionDeArrays, numeroA)
			}
		}
	}

	return interseccionDeArrays
}
