package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/scttfrdmn/ark/internal/training"
)

// handleCheckPolicy evaluates training gate policies for an action
func handleCheckPolicy(trainingSvc *training.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID          string                 `json:"user_id"`
			Action          string                 `json:"action"`
			ResourceType    string                 `json:"resource_type"`
			ResourceDetails map[string]interface{} `json:"resource_details"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			slog.Error("failed to decode policy check request", "error", err)
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "Invalid request body",
			})
			return
		}

		// Validate required fields
		if req.UserID == "" || req.Action == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "user_id and action are required",
			})
			return
		}

		// Check training gate
		decision, err := trainingSvc.CheckTrainingGate(r.Context(), req.UserID, req.Action)
		if err != nil {
			slog.Error("failed to check training gate",
				"error", err,
				"user_id", req.UserID,
				"action", req.Action,
			)
			writeJSON(w, http.StatusInternalServerError, map[string]string{
				"error": "Failed to evaluate policy",
			})
			return
		}

		slog.Info("policy evaluated",
			"user_id", req.UserID,
			"action", req.Action,
			"decision", decision.Action,
		)

		writeJSON(w, http.StatusOK, decision)
	}
}

// handleGetUserProgress retrieves training progress for a user
func handleGetUserProgress(trainingSvc *training.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "user_id")
		if userID == "" {
			writeJSON(w, http.StatusBadRequest, map[string]string{
				"error": "user_id parameter required",
			})
			return
		}

		progress, err := trainingSvc.GetUserProgress(r.Context(), userID)
		if err != nil {
			slog.Error("failed to get user progress",
				"error", err,
				"user_id", userID,
			)
			writeJSON(w, http.StatusInternalServerError, map[string]string{
				"error": "Failed to retrieve training progress",
			})
			return
		}

		writeJSON(w, http.StatusOK, map[string]interface{}{
			"user_id":  userID,
			"progress": progress,
		})
	}
}
