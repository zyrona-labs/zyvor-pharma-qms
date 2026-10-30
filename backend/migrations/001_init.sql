-- 001_init.sql
-- Core tables for Zyvor Pharma QMS
-- Append-only audit_events; no updates or deletes on audit.

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

-- Append-only audit trail. Never UPDATE or DELETE from this table.
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

-- Prevent updates/deletes on audit (defense in depth)
CREATE OR REPLACE RULE audit_no_update AS ON UPDATE TO audit_events DO INSTEAD NOTHING;
CREATE OR REPLACE RULE audit_no_delete AS ON DELETE TO audit_events DO INSTEAD NOTHING;
