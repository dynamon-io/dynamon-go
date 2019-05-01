package dynamon

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Line struct {
	Title  string  `json:"title"`
	Points []Point `json:"points"`
}

// PlotLine plots a line.
func PlotLine(path string, xs, ys []float64, opts Opts) {
	points := make([]Point, len(xs))
	for i := range xs {
		points[i] = Point{
			X: xs[i],
			Y: ys[i],
		}
	}
	api(path, payload{
		Widgets: []widget{
			widget{
				Type:   "line",
				Title:  opts.Title,
				XLabel: opts.XLabel,
				YLabel: opts.YLabel,
				Lines: []Line{
					Line{
						Title:  opts.Legend,
						Points: points,
					},
				},
			},
		},
	})
}

// PlotLines plots multiple lines.
func PlotLines(path string, xs []float64, yss [][]float64, opts Opts) {
	lines := make([]Line, len(yss))
	for i := range lines {
		points := make([]Point, len(xs))
		for j := range points {
			points[j] = Point{X: xs[j], Y: yss[i][j]}
		}

		legend := ""
		if len(opts.Legends) > i {
			legend = opts.Legends[i]
		}

		line := Line{
			Title:  legend,
			Points: points,
		}
		lines[i] = line
	}

	api(path, payload{
		Widgets: []widget{
			widget{
				Type:   "line",
				Title:  opts.Title,
				XLabel: opts.XLabel,
				YLabel: opts.YLabel,
				Lines:  lines,
			},
		},
	})
}

// PlotLinesStructured plots multiple lines from already structured line objects.
func PlotLinesStructured(path string, lines []Line, opts Opts) {
	api(path, payload{
		Widgets: []widget{
			widget{
				Type:   "line",
				Title:  opts.Title,
				XLabel: opts.XLabel,
				YLabel: opts.YLabel,
				Lines:  lines,
			},
		},
	})
}
