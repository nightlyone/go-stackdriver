package stackdriver

// A custom metric to be sent to custom metric endpoint
type Metrics struct {
	Timestamp    int64       `json:"timestamp"`
	ProtoVersion int         `json:"proto_version"`
	Data         []Datapoint `json:"data"`
}

func (m *Metrics) Submit() error {
	return submit(m)
}

// A datapoint of a custom metric
type Datapoint struct {
	Name        string  `json:"name"`
	Value       float64 `json:"value"`
	CollectedAt int64   `json:"collected_at"`
}
