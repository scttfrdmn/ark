package main

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/scttfrdmn/ark/internal/agent/aws"
)

// handleCreateBucket handles S3 bucket creation requests
func (s *server) handleCreateBucket(w http.ResponseWriter, r *http.Request) {
	var req struct {
		BucketName        string `json:"bucket_name"`
		Region            string `json:"region"`
		Encryption        struct {
			Type     string `json:"type"`
			KMSKeyID string `json:"kms_key_id,omitempty"`
		} `json:"encryption"`
		VersioningEnabled bool   `json:"versioning_enabled"`
		Profile           string `json:"profile"`
	}

	// Parse request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("failed to decode request", "error", err)
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
		return
	}

	// Validate required fields
	if req.BucketName == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "bucket_name is required",
		})
		return
	}

	// Default profile
	if req.Profile == "" {
		req.Profile = "default"
	}

	// Get credentials from store
	creds, err := s.store.GetCredential(req.Profile)
	if err != nil {
		slog.Error("failed to get credentials", "error", err, "profile", req.Profile)
		writeJSON(w, http.StatusNotFound, map[string]string{
			"error": "Credentials not found for profile: " + req.Profile,
		})
		return
	}

	// Create AWS client
	client, err := aws.NewClientFromCredentials(r.Context(), creds, req.Region)
	if err != nil {
		slog.Error("failed to create AWS client", "error", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": "Failed to initialize AWS client",
		})
		return
	}

	// Default encryption type
	if req.Encryption.Type == "" {
		req.Encryption.Type = "AES256"
	}

	// Execute S3 CreateBucket
	slog.Info("creating S3 bucket",
		"bucket", req.BucketName,
		"region", req.Region,
		"encryption", req.Encryption.Type,
		"versioning", req.VersioningEnabled,
	)

	output, err := aws.CreateBucket(r.Context(), client, aws.CreateBucketInput{
		BucketName:        req.BucketName,
		Region:            req.Region,
		EncryptionType:    req.Encryption.Type,
		KMSKeyID:          req.Encryption.KMSKeyID,
		VersioningEnabled: req.VersioningEnabled,
	})

	if err != nil {
		slog.Error("failed to create bucket",
			"error", err,
			"bucket", req.BucketName,
		)
		writeJSON(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	slog.Info("bucket created successfully",
		"bucket", output.BucketName,
		"region", output.Region,
		"location", output.Location,
	)

	// TODO Phase 5: Send audit log to backend (non-blocking)
	// go s.sendAuditLog(...)

	writeJSON(w, http.StatusCreated, output)
}
