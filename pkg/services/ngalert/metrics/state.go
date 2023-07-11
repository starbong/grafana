package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type State struct {
	AlertState          *prometheus.GaugeVec
	StateUpdateDuration *prometheus.HistogramVec
	StateUpdateCount    *prometheus.CounterVec
}

func NewStateMetrics(r prometheus.Registerer) *State {
	return &State{
		AlertState: promauto.With(r).NewGaugeVec(prometheus.GaugeOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      "alerts",
			Help:      "How many alerts by state.",
		}, []string{"state"}),
		StateUpdateDuration: promauto.With(r).NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: Namespace,
				Subsystem: Subsystem,
				Name:      "state_calculation_duration_milliseconds",
				Help:      "The duration of calculation of a single state",
				Buckets:   []float64{10, 25, 50, 100, 500, 1000, 10000, 100000},
			},
			[]string{"needsImage"},
		),
		StateUpdateCount: promauto.With(r).NewCounterVec(
			prometheus.CounterOpts{
				Namespace: Namespace,
				Subsystem: Subsystem,
				Name:      "state_calculation_total",
				Help:      "Total number of state calculations",
			},
			[]string{"needsImage"},
		),
	}
}
