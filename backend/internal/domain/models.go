package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

// Batch is the central production record.
type Batch struct {
	ID              string          `json:"id"`
	Number          string          `json:"number"`
	ProductName     string          `json:"product_name"`
	RecipeID        string          `json:"recipe_id"`
	Status          BatchStatus     `json:"status"`
	PlannedQuantity decimal.Decimal `json:"planned_quantity"`
	ActualYield     decimal.Decimal `json:"actual_yield"`
	StartedAt       *time.Time      `json:"started_at,omitempty"`
	CompletedAt     *time.Time      `json:"completed_at,omitempty"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type BatchStatus string

const (
	BatchStatusDraft      BatchStatus = "draft"
	BatchStatusInProgress BatchStatus = "in_progress"
	BatchStatusLab        BatchStatus = "laboratory"
	BatchStatusQA         BatchStatus = "quality_review"
	BatchStatusApproved   BatchStatus = "approved"
	BatchStatusReleased   BatchStatus = "released"
	BatchStatusRecalled   BatchStatus = "recalled"
	BatchStatusRejected   BatchStatus = "rejected"
)

// Sample represents a laboratory sample taken from a batch.
type Sample struct {
	ID          string    `json:"id"`
	BatchID     string    `json:"batch_id"`
	SampleCode  string    `json:"sample_code"`
	TakenAt     time.Time `json:"taken_at"`
	TakenBy     string    `json:"taken_by"`
	Status      string    `json:"status"` // pending, testing, completed, rejected
	CreatedAt   time.Time `json:"created_at"`
}

// TestResult is an individual analytical result against a specification.
type TestResult struct {
	ID          string          `json:"id"`
	SampleID    string          `json:"sample_id"`
	TestName    string          `json:"test_name"`
	Specification string        `json:"specification"`
	Result      string          `json:"result"`
	Unit        string          `json:"unit"`
	Pass        bool            `json:"pass"`
	TestedBy    string          `json:"tested_by"`
	TestedAt    time.Time       `json:"tested_at"`
	CreatedAt   time.Time       `json:"created_at"`
}

// AuditEvent is an append-only record of every significant action.
type AuditEvent struct {
	ID          string    `json:"id"`
	ActorID     string    `json:"actor_id"`
	ActorName   string    `json:"actor_name"`
	Action      string    `json:"action"`
	Resource    string    `json:"resource"`
	ResourceID  string    `json:"resource_id"`
	Details     string    `json:"details"` // JSON or text
	OccurredAt  time.Time `json:"occurred_at"`
	IPAddress   string    `json:"ip_address,omitempty"`
}
