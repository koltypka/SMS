package system

import (
	"github.com/koltypka/SMS/kolRand"
	"github.com/koltypka/SMS/visitor"
)

type Counter struct {
	costs  uint16 //издержки для обслуживания канала
	profit uint16
}

//1-5 каналов
//каждые 7 минут показательно прибывают покупатели
//каждые 9 минут показательно обслуживаем
//в среднем покупатели покупают на сумму 500 рублей
//время функционирования системы 8 часов
func StratSystem() {

}

func MakeVisitors() []visitor.Visitor {
	kRnd := kolRand.New()
	arVisitor := []visitor.Visitor{}

	timer := uint16(0)

	for timer < 480 {
		curVisitor := visitor.New(
			kRnd.MakeExp(7),
			kRnd.MakeExp(9),
			kRnd.MakeUniform(1000))

		arVisitor = append(arVisitor, curVisitor)
		timer += curVisitor.TimeTo.Visit
	}

	return arVisitor
}
