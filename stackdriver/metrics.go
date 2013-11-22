package stackdriver

// The Metrics defined by you and to be sent to custom metric endpoint
type Metrics struct {
	Timestamp    int64       `json:"timestamp"`
	ProtoVersion int         `json:"proto_version"`
	Data         []Datapoint `json:"data"`
}

// Submit the metrics to stackdriver.com
func (m *Metrics) Submit() error {
	return submit(m)
}

// A Datapoint of a Metric
type Datapoint struct {
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	CollectedAt int64   `json:"collected_at"`
}
