package visitor

type TimeTo struct {
	Visit   float64 //через какое время после предыдущего посетитель пришёл
	Service float64 //время обслуживания
}

type Visitor struct {
	TimeTo         //время нахождения в очереди
	Pay    float64 //сколько заплатил посетитель
}

func New(visit, service, pay float64) Visitor {
	return Visitor{TimeTo{visit, service}, pay}
}
