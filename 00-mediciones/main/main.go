package main

import (
	"busquedas"
	"fmt"
	"sort"
	"time"
	"utiles"
)

func main() {
	arreglo := utiles.GenerarArreglo(10, 10)
	buscado := 5

	//fmt.Println(arreglo)

	inicio := time.Now()
	// Busqueda Lineal
	fmt.Println(busquedas.BusLineal(arreglo, buscado))
	fmt.Println("Busqueda Lineal: ", time.Since(inicio))

	inicio = time.Now()
	// Ordenar el arreglo para la busqueda binaria
	sort.Ints(arreglo)
	fmt.Println("Ordenamiento: ", time.Since(inicio))
	//fmt.Println(arreglo)

	inicio = time.Now()
	// Busqueda Binaria
	fmt.Println(busquedas.BusquedaBinaria(arreglo, buscado))
	fmt.Println("Busqueda Binaria: ", time.Since(inicio))

	inicio = time.Now()
	// Ordenamiento por Burbujeo
	busquedas.OrdBurbujeo(arreglo)
	fmt.Println("Ordenamiento Burbujeo: ", time.Since(inicio))
}
