package statsd

import "github.com/go-kit/kit/metrics2"

var (
	_ metrics.Counter   = (*Counter)(nil)
	_ metrics.Gauge     = (*Gauge)(nil)
	_ metrics.Histogram = (*Timer)(nil)
	_ metrics.Histogram = (*Histogram)(nil)
)