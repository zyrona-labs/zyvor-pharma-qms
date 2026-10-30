package store

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
	"github.com/zyrona-labs/zyvor-pharma-qms/backend/internal/domain"
)

func (db *DB) CreateBatch(ctx context.Context, b *domain.Batch, actorID, actorName, ip string) error {
	err := db.Pool.QueryRow(ctx, `
		INSERT INTO batches (number, product_name, recipe_id, status, planned_quantity, actual_yield)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`, b.Number, b.ProductName, b.RecipeID, b.Status, b.PlannedQuantity, b.ActualYield).
		Scan(&b.ID, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return fmt.Errorf("insert batch: %w", err)
	}

	return db.WriteAudit(ctx, &domain.AuditEvent{
		ActorID:    actorID,
		ActorName:  actorName,
		Action:     "create",
		Resource:   "batch",
		ResourceID: b.ID,
		Details:    fmt.Sprintf("created batch %s (%s)", b.Number, b.ProductName),
		IPAddress:  ip,
	})
}

func (db *DB) ListBatches(ctx context.Context) ([]domain.Batch, error) {
	rows, err := db.Pool.Query(ctx, `
		SELECT id, number, product_name, recipe_id, status,
		       planned_quantity, actual_yield, started_at, completed_at,
		       created_at, updated_at
		FROM batches
		ORDER BY created_at DESC
		LIMIT 200
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []domain.Batch
	for rows.Next() {
		var b domain.Batch
		var planned, actual *decimal.Decimal
		err := rows.Scan(
			&b.ID, &b.Number, &b.ProductName, &b.RecipeID, &b.Status,
			&planned, &actual, &b.StartedAt, &b.CompletedAt,
			&b.CreatedAt, &b.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		if planned != nil {
			b.PlannedQuantity = *planned
		}
		if actual != nil {
			b.ActualYield = *actual
		}
		out = append(out, b)
	}
	return out, rows.Err()
}

func (db *DB) GetBatch(ctx context.Context, id string) (*domain.Batch, error) {
	var b domain.Batch
	var planned, actual *decimal.Decimal
	err := db.Pool.QueryRow(ctx, `
		SELECT id, number, product_name, recipe_id, status,
		       planned_quantity, actual_yield, started_at, completed_at,
		       created_at, updated_at
		FROM batches WHERE id = $1
	`, id).Scan(
		&b.ID, &b.Number, &b.ProductName, &b.RecipeID, &b.Status,
		&planned, &actual, &b.StartedAt, &b.CompletedAt,
		&b.CreatedAt, &b.UpdatedAt,
	)
	if err == pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if planned != nil {
		b.PlannedQuantity = *planned
	}
	if actual != nil {
		b.ActualYield = *actual
	}
	return &b, nil
}

func (db *DB) WriteAudit(ctx context.Context, e *domain.AuditEvent) error {
	if e.OccurredAt.IsZero() {
		e.OccurredAt = time.Now()
	}
	_, err := db.Pool.Exec(ctx, `
		INSERT INTO audit_events (actor_id, actor_name, action, resource, resource_id, details, occurred_at, ip_address)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`, e.ActorID, e.ActorName, e.Action, e.Resource, e.ResourceID, e.Details, e.OccurredAt, e.IPAddress)
	return err
}

func (db *DB) Migrate(ctx context.Context) error {
	// Simple embedded migration for early development.
	// Replace with golang-migrate or similar for production.
	sql := `
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS batches (
    id              UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    number          TEXT NOT NULL UNIQUE,
    product_name    TEXT NOT NULL,
    recipe_id       TEXT,
    status          TEXT NOT NULL DEFAULT 'draft',
    planned_quantity NUMERIC(18,4),
    actual_yield    NUMERIC(18,4),
    started_at      TIMESTAMPTZ,
    completed_at    TIMESTAMPTZ,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_batches_status ON batches(status);
CREATE INDEX IF NOT EXISTS idx_batches_number ON batches(number);

CREATE TABLE IF NOT EXISTS samples (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    batch_id    UUID NOT NULL REFERENCES batches(id),
    sample_code TEXT NOT NULL,
    taken_at    TIMESTAMPTZ NOT NULL,
    taken_by    TEXT NOT NULL,
    status      TEXT NOT NULL DEFAULT 'pending',
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_samples_batch ON samples(batch_id);

CREATE TABLE IF NOT EXISTS test_results (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    sample_id     UUID NOT NULL REFERENCES samples(id),
    test_name     TEXT NOT NULL,
    specification TEXT,
    result        TEXT,
    unit          TEXT,
    pass          BOOLEAN,
    tested_by     TEXT,
    tested_at     TIMESTAMPTZ,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_test_results_sample ON test_results(sample_id);

CREATE TABLE IF NOT EXISTS audit_events (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    actor_id    TEXT NOT NULL,
    actor_name  TEXT NOT NULL,
    action      TEXT NOT NULL,
    resource    TEXT NOT NULL,
    resource_id TEXT,
    details     TEXT,
    occurred_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    ip_address  TEXT
);

CREATE INDEX IF NOT EXISTS idx_audit_resource ON audit_events(resource, resource_id);
CREATE INDEX IF NOT EXISTS idx_audit_occurred ON audit_events(occurred_at DESC);

CREATE OR REPLACE RULE audit_no_update AS ON UPDATE TO audit_events DO INSTEAD NOTHING;
CREATE OR REPLACE RULE audit_no_delete AS ON DELETE TO audit_events DO INSTEAD NOTHING;
`
	_, err := db.Pool.Exec(ctx, sql)
	return err
}
