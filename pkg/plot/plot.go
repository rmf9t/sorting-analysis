package plot

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"msis/pkg/models"
)

func PlotGraph(ins []models.SortTime, ms []models.SortTime, filename string) error {
	p := plot.New()

	p.Title.Text = "Insertion and Merge Sort Analysis"
	p.X.Label.Text = "Array Size"
	p.Y.Label.Text = "Time (nanoseconds)"

	err := plotutil.AddLinePoints(p,
		"Insertion Sort", addPoints(ins),
		"Merge Sort", addPoints(ms))
	if err != nil {
		return err
	}

	if err := p.Save(10*vg.Inch, 10*vg.Inch, filename); err != nil {
		return err
	}
	return nil
}

func addPoints(n []models.SortTime) plotter.XYs {
	pts := make(plotter.XYs, len(n))
	for i := range pts {
		pts[i].X = float64(n[i].Size)
		pts[i].Y = float64(n[i].TimeTaken)
	}
	return pts
}
