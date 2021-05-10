package core

type DomainValidationFailed struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type DomainEntityNotFound struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}
