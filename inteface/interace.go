package inteface

type Interface struct {
	Revenue   []int
	Effective []int
	N         int32
}

func NewInterface() Interface {
	return Interface{}
}

func (i *Interface) AddResult(Revenue, Effective int) {
	i.Revenue = append(i.Revenue, Revenue)
	i.Effective = append(i.Effective, Effective)
	i.N++
}

func (i *Interface) ScoreResult() (int, int) {
	return i.scoreMean(i.Revenue), i.scoreMean(i.Effective)
}

func (i *Interface) scoreMean(slice []int) int {
	var intResult int = 0

	for _, value := range slice {
		intResult += int(value)
	}

	return intResult / int(i.N)
}
