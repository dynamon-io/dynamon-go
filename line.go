package dynamon

// LineOpts specifies options for Line
type LineOpts struct {
	Title  string
	XLabel string
	YLabel string
	Legend string
}

// Line plots a line
func Line(path string, xs, ys []float64, opts LineOpts) {
	points := make([]point, len(xs))
	for i := range xs {
		points[i] = point{
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
				Lines: []line{
					line{
						Title:  opts.Legend,
						Points: points,
					},
				},
			},
		},
	})
}
