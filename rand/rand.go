package rand

import (
	"math/rand"
	"time"
)

//1-5 каналов
//каждые 7 минут показательно прибывают покупатели
//каждые 9 минут показательно обслуживаем
//в среднем покупатели покупают на сумму 500 рублей
//время функционирования системы 8 часов
func Test() {

}

//генератор случайных чисел работает только внутри тела функции
func MakeRavnr() float64 {
	//задаём семя
	generator := rand.New(rand.NewSource(time.Now().UnixNano()))
	//генерируем значение
	returnFloat64 := generator.ExpFloat64()

	return returnFloat64
}

func main() {

}
