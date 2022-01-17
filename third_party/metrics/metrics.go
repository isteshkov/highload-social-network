package external

import "github.com/prometheus/client_golang/prometheus"

type Metrics interface {
	NotifyRequestDone(path string, latency float64)
}

func NewMetrics() (*metrics, error) {
	requestsMetric := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace: "somerset",
		Name:      "api_calls",
		Help:      "latency groups by paths in seconds",
	}, []string{"path"})

	err := prometheus.Register(requestsMetric)
	if err != nil {
		return nil, err
	}

	return &metrics{
		requestsMetric: requestsMetric,
	}, nil
}

type metrics struct {
	requestsMetric *prometheus.SummaryVec
}

func (m *metrics) NotifyRequestDone(path string, latency float64) {
	m.requestsMetric.WithLabelValues(path).Observe(latency)
}

func NewMetricsMock() Metrics {
	return &metricsMock{}
}

type metricsMock struct {
}

func (m metricsMock) NotifyRequestDone(path string, latency float64) {
	return
}
