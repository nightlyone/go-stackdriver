package stackdriver

// an annotion made by someone
type Annotation struct {
	Message     string `json:"message"`
	AnnotatedBy string `json:"annotated_by,omitempty"`
	Level       string `json:"level,omitempty"`
	IntanceId   string `json:"instance_id,omitempty"`
	EventEpoch  int64  `json:"event_epoch,omitempty"`
}

func (a *Annotation) Submit() error {
	return submit(a)
}
