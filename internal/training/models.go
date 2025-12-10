package training

// Module represents a training module
type Module struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Title            string `json:"title"`
	EstimatedMinutes int    `json:"estimated_minutes"`
}

// PolicyDecision represents the result of policy evaluation
type PolicyDecision struct {
	Action          string   `json:"action"` // "allow" or "block"
	Reason          string   `json:"reason,omitempty"`
	RequiredModules []Module `json:"required_modules,omitempty"`
	Message         string   `json:"message"`
}

// Progress represents user training progress
type Progress struct {
	ModuleID    string `json:"module_id"`
	ModuleName  string `json:"module_name"`
	Status      string `json:"status"` // not_started, in_progress, completed
	CompletedAt string `json:"completed_at,omitempty"`
}
