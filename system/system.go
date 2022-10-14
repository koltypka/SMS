package system

import (
	"fmt"

	"github.com/koltypka/SMS/visitor"
)

//1-5 каналов
//каждые 7 минут показательно прибывают покупатели
//каждые 9 минут показательно обслуживаем
//в среднем покупатели покупают на сумму 500 рублей равномерно
//время функционирования системы 8 часов
func StratSystem(numberOfCounters uint16) []visitor.Visitor {

	for i := uint16(0); i < numberOfCounters; i++ {
		Manager := NewManager()
		counter := newCounter(
			Manager.kRnd.MakeUniform(1000), //генерим затраты на кассу
			Manager.MakeVisitors())

		counter.CountVisitors()

		fmt.Println(counter.visitors)
		fmt.Println(counter.profit)
		fmt.Println(counter.costs)
		fmt.Println("============================")
	}

	Manager := NewManager()

	return Manager.MakeVisitors()
}
