package system

import (
	"github.com/koltypka/SMS/kolRand"
	"github.com/koltypka/SMS/visitor"
)

type Manager struct {
	kRnd              kolRand.KolRandom
	prevTimeToService uint16 //вычитаем время на приход для следующего посетителя
	timer             uint16 //таймер для замера моделирования СМО
}

func NewManager() Manager {
	return Manager{kolRand.New(), 0, 0}
}

// обработчик очереди
// если предыдущий посетитель ещё обрабатывается при приходе нового постетителя, то вычитаем из
// переменной обработчика очереди через какое время пришёл посетитель, и добавлям время обслуживания новой заявки
func (sysMg *Manager) handleQueue(nextTimeTo visitor.TimeTo) visitor.TimeTo {

	if sysMg.prevTimeToService >= nextTimeTo.Visit {
		sysMg.prevTimeToService = sysMg.prevTimeToService - nextTimeTo.Visit + nextTimeTo.Service
		nextTimeTo.Visit = 0
	} else {
		nextTimeTo.Visit = nextTimeTo.Visit - sysMg.prevTimeToService
		sysMg.prevTimeToService = nextTimeTo.Service
	}

	return nextTimeTo
}

func (sysMg *Manager) checkRefuse(service uint16) uint16 {
	uint16Return := service
	key := sysMg.kRnd.MakeUniform(100)

	if key > 90 {
		uint16Return = service + 5
	}

	if key > 95 {
		uint16Return = service + service
	}

	return uint16Return
}

func (sysMg *Manager) MakeVisitors() []visitor.Visitor {
	arVisitor := []visitor.Visitor{}
	sysMg.timer = 0

	for sysMg.timer+sysMg.prevTimeToService < 480 {
		curVisitor := visitor.New(
			sysMg.kRnd.MakeExp(7),
			sysMg.kRnd.MakeExp(9),
			sysMg.kRnd.MakeUniform(1000))

		curVisitor.TimeTo.Service = sysMg.checkRefuse(curVisitor.TimeTo.Service)

		sysMg.timer += curVisitor.TimeTo.Visit

		curVisitor.TimeTo = sysMg.handleQueue(curVisitor.TimeTo)

		arVisitor = append(arVisitor, curVisitor)
	}

	return arVisitor
}
