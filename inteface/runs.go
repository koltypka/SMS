package inteface

import (
	"fmt"

	"github.com/koltypka/SMS/system"
)

func makeYBars(runs int) []int {
	var i int = runs
	arReturn := []int{}
	for i > 0 {
		arReturn = append(arReturn, runs/i)
		i = i / 10
	}

	return arReturn
}

func StartRuns(number, runs int) {
	var Revenue, Effective int
	var scoreRevenue, scoreEffective int

	var arRuns []int = makeYBars(runs)

	systems := []system.System{}
	graphRevenue := NewGraph("метод 1", "Revenue")
	graphEffective := NewGraph("метод 2", "Effective")

	Interface := NewInterface()

	for i := 0; i < number; i++ {
		systems = append(systems, system.NewSystem(uint16(i))) //создаём системы с 1-5 кассами

		graphRevenue.ArLines = append(graphRevenue.ArLines, NewLine("Кол-во касс: "+fmt.Sprint(i+1)))     //создаём линии
		graphEffective.ArLines = append(graphEffective.ArLines, NewLine("Кол-во касс: "+fmt.Sprint(i+1))) //создаём линии
	}

	for i := 0; i < runs; i++ {
		for sysNumber, sys := range systems {
			Revenue, Effective = sys.StartSimulation()

			Interface.AddResult(Revenue, Effective)
			scoreRevenue, scoreEffective = Interface.ScoreResult()

			graphRevenue.ArLines[sysNumber].AddValue(scoreRevenue)
			graphEffective.ArLines[sysNumber].AddValue(scoreEffective)
		}
	}

	graphRevenue.MakeGraph(arRuns, revenue_name)
	graphEffective.MakeGraph(arRuns, effective_name)
}
