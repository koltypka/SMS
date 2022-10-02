package rand

import (
	"math/rand"
	"time"
)

//генератор случайных чисел работает только внутри тела функции
func MakeRavnr(lambda float64) uint16 {
	//задаём семя
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	//генерируем значение
	returnFloat64 := 1 / (generator.ExpFloat64() / lambda)

	return uint16(returnFloat64)

}
