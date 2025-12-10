package training

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/scttfrdmn/ark/internal/database"
)

// Service provides training policy enforcement functionality
type Service struct {
	db *database.DB
}

// NewService creates a new training service
func NewService(db *database.DB) *Service {
	return &Service{db: db}
}

// CheckTrainingGate evaluates if a user meets training requirements for an action
func (s *Service) CheckTrainingGate(ctx context.Context, userID string, action string) (*PolicyDecision, error) {
	// Query for policies that apply to this action
	query := `
		SELECT rules
		FROM policies
		WHERE policy_type = 'training_gate'
		  AND rules->>'actions' @> $1
		  AND status = 'active'
	`

	actionJSON := fmt.Sprintf(`["%s"]`, action)
	rows, err := s.db.QueryContext(ctx, query, actionJSON)
	if err != nil {
		return nil, fmt.Errorf("query policies: %w", err)
	}
	defer rows.Close()

	// Collect required modules from all matching policies
	requiredModules := make(map[string]bool)
	for rows.Next() {
		var rulesJSON []byte
		if err := rows.Scan(&rulesJSON); err != nil {
			return nil, fmt.Errorf("scan rules: %w", err)
		}

		var rules struct {
			RequiredModules []string `json:"required_modules"`
		}
		if err := json.Unmarshal(rulesJSON, &rules); err != nil {
			return nil, fmt.Errorf("unmarshal rules: %w", err)
		}

		for _, moduleName := range rules.RequiredModules {
			requiredModules[moduleName] = true
		}
	}

	// If no policies match, allow the action
	if len(requiredModules) == 0 {
		return &PolicyDecision{
			Action:  "allow",
			Message: "No training requirements for this action",
		}, nil
	}

	// Check if user has completed all required modules
	moduleNames := make([]string, 0, len(requiredModules))
	for name := range requiredModules {
		moduleNames = append(moduleNames, name)
	}

	// Query for incomplete modules
	incompleteQuery := `
		SELECT tm.id, tm.name, tm.title, tm.estimated_minutes
		FROM training_modules tm
		LEFT JOIN user_training_progress utp
			ON tm.id = utp.module_id AND utp.user_id = $1 AND utp.status = 'completed'
		WHERE tm.name = ANY($2)
		  AND utp.id IS NULL
	`

	incompleteRows, err := s.db.QueryContext(ctx, incompleteQuery, userID, moduleNames)
	if err != nil {
		return nil, fmt.Errorf("query incomplete modules: %w", err)
	}
	defer incompleteRows.Close()

	var incomplete []Module
	for incompleteRows.Next() {
		var module Module
		if err := incompleteRows.Scan(&module.ID, &module.Name, &module.Title, &module.EstimatedMinutes); err != nil {
			return nil, fmt.Errorf("scan module: %w", err)
		}
		incomplete = append(incomplete, module)
	}

	// If any modules are incomplete, block the action
	if len(incomplete) > 0 {
		return &PolicyDecision{
			Action:          "block",
			Reason:          "training_required",
			RequiredModules: incomplete,
			Message:         "Complete required training modules to perform this operation",
		}, nil
	}

	// All required modules completed - allow
	return &PolicyDecision{
		Action:  "allow",
		Message: "Training requirements met",
	}, nil
}

// GetUserProgress retrieves training progress for a user
func (s *Service) GetUserProgress(ctx context.Context, userID string) ([]Progress, error) {
	query := `
		SELECT
			tm.id,
			tm.name,
			COALESCE(utp.status, 'not_started') as status,
			utp.completed_at
		FROM training_modules tm
		LEFT JOIN user_training_progress utp
			ON tm.id = utp.module_id AND utp.user_id = $1
		ORDER BY tm.name
	`

	rows, err := s.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, fmt.Errorf("query progress: %w", err)
	}
	defer rows.Close()

	var progress []Progress
	for rows.Next() {
		var p Progress
		var completedAt sql.NullTime

		if err := rows.Scan(&p.ModuleID, &p.ModuleName, &p.Status, &completedAt); err != nil {
			return nil, fmt.Errorf("scan progress: %w", err)
		}

		if completedAt.Valid {
			p.CompletedAt = completedAt.Time.Format("2006-01-02T15:04:05Z")
		}

		progress = append(progress, p)
	}

	return progress, nil
}
