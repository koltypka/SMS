package system

const mean_costs int32 = 2000

type System struct {
	numberOfCounters uint16 //число касс в СМО
	totalCosts       uint16
	totalRevenues    uint16
	totalVisitors    uint16
}

func (sys *System) totalSumm(counter Counter) {
	sys.totalCosts += counter.costs
	sys.totalRevenues += counter.profit
	sys.totalVisitors += counter.visitors
}

func NewSystem(numberOfCounters uint16) System {
	return System{numberOfCounters, 0, 0, 0}
}

func (sys *System) ResultRevenue() int32 {
	return int32(sys.totalRevenues) - int32(sys.totalCosts)
}

func (sys *System) ResultEffective() int32 {
	return int32(float32(float32(sys.totalCosts)/float32(sys.totalRevenues+1)) * 100)
}

//1-5 каналов
//каждые 7 минут показательно прибывают покупатели
//каждые 9 минут показательно обслуживаем
//в среднем покупатели покупают на сумму 500 рублей равномерно
//время функционирования системы 8 часов
func (sys *System) StartSimulation() (int, int) {

	for i := uint16(0); i < sys.numberOfCounters; i++ {
		Manager := NewManager()
		counter := NewCounter(
			Manager.kRnd.MakeUniform(mean_costs), //генерим затраты на кассу
			Manager.makeVisitors())

		counter.prepareResult()

		sys.totalSumm(counter)
	}

	return int(sys.ResultRevenue()), int(sys.ResultEffective())
}
