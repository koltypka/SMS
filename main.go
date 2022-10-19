package main

import (
	"github.com/koltypka/SMS/inteface"
	"github.com/koltypka/SMS/system"
)

func main() {
	var Revenue, Effective int
	var scoreRevenue, scoreEffective int
	number := 5
	runs := 10000

	systems := []system.System{}
	graphRevenue := inteface.NewGraph("метод 1", "Revenue")
	graphEffective := inteface.NewGraph("метод 2", "Effective")
	//graph := inteface.NewGraph("тест работы", "тест")
	Interface := inteface.NewInterface()

	for i := 0; i < number; i++ {
		systems = append(systems, system.NewSystem(uint16(i)))                                         //создаём системы с 1-5 кассами
		graphRevenue.ArLines = append(graphRevenue.ArLines, inteface.NewLine("Граффик"+string(i)))     //создаём линии
		graphEffective.ArLines = append(graphEffective.ArLines, inteface.NewLine("Граффик"+string(i))) //создаём линии
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
	graphRevenue.MakeGraph([]int{1, 10, 100, 1000, 10000}, "revenue_name.html")
	graphEffective.MakeGraph([]int{1, 10, 100, 1000, 10000}, "effective.html")
}
