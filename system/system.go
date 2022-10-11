package system

import (
	"github.com/koltypka/SMS/kolRand"
	"github.com/koltypka/SMS/visitor"
)

type SystemManager struct {
	kRnd              kolRand.KolRandom
	prevTimeToService uint16 //вычитаем время на приход для следующего посетителя
	timer             uint16 //таймер для замера моделирования СМО
}

type Counter struct {
	costs  uint16 //издержки для обслуживания канала
	profit uint16
}

func checkRefuse(key, service uint16) uint16 {
	uint16Return := service

	if key > 90 {
		uint16Return = service + 5
	}

	if key > 95 {
		uint16Return = service + service
	}

	return uint16Return
}

//обработчик очереди
func (sysMg *SystemManager) handleQueue(nextTimeTo visitor.TimeTo) visitor.TimeTo {

	if sysMg.prevTimeToService >= nextTimeTo.Visit {
		sysMg.prevTimeToService = sysMg.prevTimeToService - nextTimeTo.Visit
		nextTimeTo.Visit = 0
	} else {
		nextTimeTo.Visit = nextTimeTo.Visit - sysMg.prevTimeToService
		sysMg.prevTimeToService = nextTimeTo.Service
	}

	return nextTimeTo
}

func (sysMg *SystemManager) MakeVisitors() []visitor.Visitor {
	arVisitor := []visitor.Visitor{}
	sysMg.timer = 0

	for sysMg.timer < 480 {
		curVisitor := visitor.New(
			sysMg.kRnd.MakeExp(7),
			sysMg.kRnd.MakeExp(9),
			sysMg.kRnd.MakeUniform(1000))

		curVisitor.TimeTo.Service = checkRefuse(
			sysMg.kRnd.MakeUniform(100),
			curVisitor.TimeTo.Service)

		curVisitor.TimeTo = sysMg.handleQueue(curVisitor.TimeTo)

		arVisitor = append(arVisitor, curVisitor)
		sysMg.timer += curVisitor.TimeTo.Visit
	}

	return arVisitor
}

func New() SystemManager {
	return SystemManager{kolRand.New(), 0, 0}
}

//1-5 каналов
//каждые 7 минут показательно прибывают покупатели
//каждые 9 минут показательно обслуживаем
//в среднем покупатели покупают на сумму 500 рублей равномерно
//время функционирования системы 8 часов
func StratSystem() []visitor.Visitor {
	systemManager := New()

	return systemManager.MakeVisitors()
}
