package system

import (
	"github.com/koltypka/SMS/kolRand"
	"github.com/koltypka/SMS/visitor"
)

//1-5 каналов
//каждые 7 минут показательно прибывают покупатели
//каждые 9 минут показательно обслуживаем
//в среднем покупатели покупают на сумму 500 рублей
//время функционирования системы 8 часов
func StratSystem() {

}

func MakeVisitors() visitor.Visitor {
	kRnd := kolRand.New()
	test := visitor.New(
		kRnd.MakeRavnr(7),
		kRnd.MakeRavnr(9),
		kRnd.MakeRavnr(500))

	return test
}
