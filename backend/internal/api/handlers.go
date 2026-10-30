package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"log/slog"

	"github.com/zyrona-labs/zyvor-pharma-qms/backend/internal/domain"
	"github.com/zyrona-labs/zyvor-pharma-qms/backend/internal/store"
)

func RegisterRoutes(r chi.Router, db *store.DB, logger *slog.Logger) {
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/batches", listBatches(db, logger))
		r.Post("/batches", createBatch(db, logger))
		r.Get("/batches/{id}", getBatch(db, logger))
		// TODO: lab samples, test results, approvals, recalls, audit export
	})
}

func listBatches(db *store.DB, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Placeholder — replace with real query
		batches := []domain.Batch{}
		writeJSON(w, http.StatusOK, batches)
	}
}

func createBatch(db *store.DB, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var b domain.Batch
		if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		// TODO: validate, persist, write audit event
		b.Status = domain.BatchStatusDraft
		writeJSON(w, http.StatusCreated, b)
	}
}

func getBatch(db *store.DB, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		// TODO: load from DB
		writeJSON(w, http.StatusOK, map[string]string{"id": id, "status": "placeholder"})
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
