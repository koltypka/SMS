package system

import (
	"github.com/koltypka/SMS/visitor"
	"github.com/koltypka/kolRand"
)

const visit_time_mean float64 = 7
const service_time_mean float64 = 9
const mean_profit int32 = 1000
const first_type_refusal_chance uint16 = 90
const second_type_refusal_chance uint16 = 95

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

	if key > first_type_refusal_chance {
		uint16Return = service + 5
	}

	if key > second_type_refusal_chance {
		uint16Return = service + service
	}

	return uint16Return
}

func (sysMg *Manager) makeVisitors() []visitor.Visitor {
	arVisitor := []visitor.Visitor{}
	sysMg.timer = 0

	for sysMg.timer+sysMg.prevTimeToService < 480 {
		curVisitor := visitor.NewVisitor(
			sysMg.kRnd.MakeExp(visit_time_mean),
			sysMg.kRnd.MakeExp(service_time_mean),
			sysMg.kRnd.MakeUniform(mean_profit))

		curVisitor.TimeTo.Service = sysMg.checkRefuse(curVisitor.TimeTo.Service)

		sysMg.timer += curVisitor.TimeTo.Visit

		curVisitor.TimeTo = sysMg.handleQueue(curVisitor.TimeTo)

		arVisitor = append(arVisitor, curVisitor)
	}

	return arVisitor
}
