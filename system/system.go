package system

import (
	"github.com/koltypka/SMS/visitor"
)

//1-5 каналов
//каждые 7 минут показательно прибывают покупатели
//каждые 9 минут показательно обслуживаем
//в среднем покупатели покупают на сумму 500 рублей равномерно
//время функционирования системы 8 часов
func StratSystem() []visitor.Visitor {
	Manager := NewManager()

	return Manager.MakeVisitors()
}
