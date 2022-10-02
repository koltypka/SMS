package visitor

type TimeTo struct {
	Visit   uint16 //через какое время после предыдущего посетитель пришёл
	Service uint16 //время обслуживания
}

type Visitor struct {
	TimeTo        //время нахождения в очереди
	Pay    uint16 //сколько заплатил посетитель
}

func New(visit, service, pay uint16) Visitor {
	return Visitor{TimeTo{visit, service}, pay}
}
