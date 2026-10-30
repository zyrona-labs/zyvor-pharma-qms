package api

import (
	"encoding/json"
	"net/http"
	"strings"

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
		batches, err := db.ListBatches(r.Context())
		if err != nil {
			logger.Error("list batches", "error", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		if batches == nil {
			batches = []domain.Batch{}
		}
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
		if strings.TrimSpace(b.Number) == "" || strings.TrimSpace(b.ProductName) == "" {
			http.Error(w, "number and product_name required", http.StatusBadRequest)
			return
		}
		if b.Status == "" {
			b.Status = domain.BatchStatusDraft
		}

		// Temporary actor until real auth is wired
		actorID := r.Header.Get("X-Actor-ID")
		if actorID == "" {
			actorID = "system"
		}
		actorName := r.Header.Get("X-Actor-Name")
		if actorName == "" {
			actorName = "System"
		}
		ip := r.RemoteAddr

		if err := db.CreateBatch(r.Context(), &b, actorID, actorName, ip); err != nil {
			logger.Error("create batch", "error", err)
			http.Error(w, "failed to create batch", http.StatusInternalServerError)
			return
		}
		writeJSON(w, http.StatusCreated, b)
	}
}

func getBatch(db *store.DB, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		b, err := db.GetBatch(r.Context(), id)
		if err != nil {
			logger.Error("get batch", "error", err, "id", id)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		if b == nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		writeJSON(w, http.StatusOK, b)
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
