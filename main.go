package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func burbuja(b []int, p int) {
	var temp, i, j int

	for i = 0; i <= p; i++ {
		for j = 0; j < p-1; j++ {
			if b[j] > b[j+1] {
				temp = b[j]
				b[j] = b[j+1]
				b[j+1] = temp
			}
		}
	}
}

func insercionDirecta(b []int) {
	var i, j, k int

	for j = 1; j < len(b); j++ {
		i = j - 1
		k = b[j]
		for i >= 0 {
			if k < b[i] {
				b[i+1] = b[i]
				i = i - 1
			} else {
				break
			}
			b[i+1] = k
		}
	}
}

func quickSort(b []int, izq int, der int) {
	var i, j, v, aux int
	i = izq
	j = der
	v = b[(izq+der)/2]

	for {
		for (b[i] < v) && (j <= der) {
			i++
		}
		for (v < b[j]) && (j > izq) {
			j--
		}
		if i <= j {
			aux = b[i]
			b[i] = b[j]
			b[j] = aux
			i++
			j--
		}
		if !(i <= j) {
			break
		}
	}
	if izq < j {
		quickSort(b, izq, j)
	}
	if izq < j {
		quickSort(b, i, der)
	}
}

/**/
func shellsort(a []int) {
	var (
		n    = len(a)
		gaps = []int{1}
		k    = 1
	)
	for {
		gap := element(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}
	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if a[j-gap] > a[j] {
					a[j-gap], a[j] = a[j], a[j-gap]
				}
				j = j - gap
			}
		}
	}
}
func element(a, b int) int {
	e := 1
	for b > 0 {
		if b&1 != 0 {
			e *= a
		}
		b >>= 1
		a *= a
	}
	return e
}

/**/

func seleccion(A []int) {
	var i, j, menor, pos, tmp int
	for i = 0; i < len(A)-1; i++ {
		menor = A[i]
		for j = i + 1; j < len(A); j++ {
			if A[j] < menor {
				menor = A[j]
				pos = j
			}
		}
		if pos != i {
			tmp = A[i]
			A[i] = A[pos]
			A[pos] = tmp
		}
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func main() {

	var option, cantidadDatos int

	fmt.Println("Bienvenido, ingrese la cantidad de datos a ordenar:")
	fmt.Scanln(&cantidadDatos)

	datos := generateSlice(cantidadDatos)

	fmt.Println("Ingresa la opción a probar:")
	fmt.Println("1)Burbuja       2)Inserción     3)ShellSort        4)QuickSort     5)Selección")
	fmt.Scanln(&option)

	switch option {
	case 1:
		fmt.Println("Burbuja:")
		start := makeTimestamp()
		burbuja(datos, len(datos))
		end := makeTimestamp()
		fmt.Printf("%d milisegundos\n\n", end-start)
		fmt.Printf("ByLuisFlahan4051! ;3\n\n")
	case 2:
		fmt.Println("InserciónDirecta:")
		start := makeTimestamp()
		insercionDirecta(datos)
		end := makeTimestamp()
		fmt.Printf("%d milisegundos\n\n", end-start)
		fmt.Printf("ByLuisFlahan4051! ;3\n\n")
	case 3:
		fmt.Println("ShellSort:")
		start := makeTimestamp()
		shellsort(datos)
		end := makeTimestamp()
		fmt.Printf("%d milisegundos\n\n", end-start)
		fmt.Printf("ByLuisFlahan4051! ;3\n\n")
	case 4:
		fmt.Println("QuickSort:")
		start := makeTimestamp()
		quickSort(datos, 0, len(datos)-1)
		end := makeTimestamp()
		fmt.Printf("%d milisegundos\n\n", end-start)
		fmt.Printf("ByLuisFlahan4051! ;3\n\n")
	case 5:
		fmt.Println("Selección:")
		start := makeTimestamp()
		seleccion(datos)
		end := makeTimestamp()
		fmt.Printf("%d milisegundos\n\n", end-start)
		fmt.Printf("ByLuisFlahan4051! ;3\n\n")
	default:
		fmt.Println("No existe la opción!")
	}

	var contin int
	fmt.Println("Preciona ENTER para finalizar...")
	fmt.Scanln(&contin)
}
