package inteface

import "github.com/go-echarts/go-echarts/v2/opts"

//содержит все значения линии
type line struct {
	values []opts.LineData
	name   string
}

func NewLine(name string) line {
	return line{[]opts.LineData{}, name}
}

//значение точек
func (l *line) AddValue(value int) {
	l.values = append(l.values, opts.LineData{Value: value})
}
