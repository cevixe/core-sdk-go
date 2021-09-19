package core

type EventHandlingFailed struct {
	Handler    string `json:"handler,omitempty"`
	Error      string `json:"error,omitempty"`
	StackTrace string `json:"stackTrace,omitempty"`
}

type DomainValidationFailed struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type DomainEntityNotFound struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
}
