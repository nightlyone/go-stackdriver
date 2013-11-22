package stackdriver

// A Deploy of some code
type Deploy struct {
	RevisionID string `json:"revision_id"`
	DeployedBy string `json:"deployed_by,omitempty"`
	DeployedTo string `json:"deployed_to,omitempty"`
	Repository string `json:"repository,omitempty"`
}

// Submit the deploy to stackdriver.com
func (d *Deploy) Submit() error {
	return submit(d)
}
