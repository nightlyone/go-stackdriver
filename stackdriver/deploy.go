package stackdriver

type Deploy struct {
	RevisionId string `json:"revision_id"`
	DeployedBy string `json:"deployed_by,omitempty"`
	DeployedTo string `json:"deployed_to,omitempty"`
	Repository string `json:"repository,omitempty"`
}
