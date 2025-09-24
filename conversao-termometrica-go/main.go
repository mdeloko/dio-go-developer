package conversaotermometricago

import(
	"fmt"
	"math"
)

func main(){
	ebuli_k := 373.15
	ebuli_c := math.Floor(ebuli_k - 273)
	fmt.Printf("A temperatura de ebulição em K é %v, e em Celsius é %v\n",ebuli_k,ebuli_c)
}