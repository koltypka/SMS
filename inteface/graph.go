package inteface

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

const effective_name = "effective.html"
const revenue_name = "revenue_name.html"

//содержит все линии со значениями
type Graph struct {
	line     *charts.Line
	ArLines  []line
	title    string
	subtitle string
}

func NewGraph(title, subtitle string) Graph {
	return Graph{charts.NewLine(), []line{}, title, subtitle}
}

//на вход имена точек
func (g *Graph) MakeGraph(arAxis []int, fileName string) {
	// set some global options like Title/Legend/ToolTip or anything else
	g.line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    g.title,
			Subtitle: g.subtitle,
		}))

	g.line.SetXAxis(arAxis)

	for _, line := range g.ArLines {
		g.line.AddSeries(line.name, line.values)
	}

	g.line.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	file, _ := os.Create(fileName)
	g.line.Render(file)
}
