package system

import "github.com/koltypka/SMS/visitor"

type Counter struct {
	costs    uint16            //издержки для обслуживания канала
	stream   []visitor.Visitor //поток заявок
	visitors uint16            //количество посетителей
	profit   uint16
}

func NewCounter(costs uint16, stream []visitor.Visitor) Counter {
	return Counter{costs, stream, 0, 0}
}

func (cnt *Counter) prepareResult() {
	for _, visitor := range cnt.stream {
		cnt.visitors++
		cnt.profit += visitor.Pay
	}
}
