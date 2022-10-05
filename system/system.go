package system

import (
	"fmt"

	"github.com/koltypka/SMS/rand"
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
	arVisitor := []visitor.Visitor{}
	var curVisitor visitor.Visitor //инициализация пустой структуры

	var i = 0
	for i > 3 {
		curVisitor := visitor.New(
			rand.MakeRavnr(7),
			rand.MakeRavnr(9),
			rand.MakeRavnr(500))

		arVisitor = append(arVisitor, curVisitor)
		fmt.Print(i)
		i++
	}

	fmt.Println(arVisitor) //вывод цикла не получается записать

	return curVisitor
}
