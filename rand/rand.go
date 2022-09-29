package rand

import "math/rand"

//1-5 каналов
//каждые 7 минут показательно прибывают покупатели
//каждые 9 минут показательно обслуживаем
//в среднем покупатели покупают на сумму 500 рублей
//время функционирования системы 8 часов
func Test() {

}

func MakeRavnr() float64 {
	rand.Seed(1)
	returnFloat64 := rand.ExpFloat64() / 7
	return returnFloat64
}

func main() {

}
