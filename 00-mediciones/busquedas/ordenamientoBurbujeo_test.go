package busquedas_test



import(
	"testing"
	"busquedas"
	"github.com/stretchr/testify/assert"
)

func TestOrdenarArregloPorBurbujeo(t *testing.T){
	arreglo := []int{-1, 5, -6, 10, 7, 2, -10, 11, 3, 17}
	
	esperado := []int{-10, -6, -1, 2, 3, 5, 7, 10, 11, 17}

	obtenido := busquedas.OrdBurbujeo(arreglo)

	assert.Equal(t, esperado, obtenido)
}

func TestOrdenarArregloNulo(t *testing.T){
	arreglo := []int{}
	
	esperado := []int{}

	obtenido := busquedas.OrdBurbujeo(arreglo)

	assert.Equal(t, esperado, obtenido)
}

func TestOrdenarArregloConDatosDuplicados(t *testing.T){
	arreglo := []int{-1, 5, -6, 10, 7, 2, -10, 5, 3, 17}
	
	esperado := []int{-10, -6, -1, 2, 3, 5, 5, 7, 10, 17}

	obtenido := busquedas.OrdBurbujeo(arreglo)

	assert.Equal(t, esperado, obtenido)
}