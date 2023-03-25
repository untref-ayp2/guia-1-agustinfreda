package busquedas

// BusBurbujeo busca un elemento en un arreglo de enteros usando el algoritmo de búsqueda lineal
func OrdBurbujeo(datos []int) []int{
	// inicio := 0
	
	for i:= 0; i < len(datos)-1 ; i++{
		for j := 0; j < len(datos)-i-1; j++{
			if(datos[j] > datos[j+1]){
				datos[j] , datos[j+1] = datos[j+1] , datos[j]
			}
		}
	}
	return datos
}