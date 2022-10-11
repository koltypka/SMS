package visitor

type Visitor struct {
	TimeTo        //время нахождения в очереди
	Pay    uint16 //сколько заплатил посетитель
}

func New(visit, service, pay uint16) Visitor {
	return Visitor{TimeTo{visit, service}, pay}
}
