package kolRand

import (
	"math/rand"
	"time"
)

type KolRandom struct {
	seed *rand.Rand //генератор случайных чисел
}

func New() KolRandom {
	return KolRandom{rand.New(rand.NewSource(time.Now().UnixNano()))}
}

//генератор случайных чисел работает только внутри тела функции
//генератор случайных чисел распределенных по показательному закону
func (rnd *KolRandom) MakeExp(lambda float64) uint16 {
	returnUnit16 := 1 / (rnd.seed.ExpFloat64() / lambda)

	return uint16(returnUnit16)

}
