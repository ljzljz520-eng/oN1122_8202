package domain

import "fmt"

type Status string

const (
	StatusDraft     Status = "draft"
	StatusReview    Status = "review"
	StatusConfirmed Status = "confirmed"
	StatusPublished Status = "published"
	StatusArchived  Status = "archived"
)

type Record struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Permission string `json:"permission"`
	Status     Status `json:"status"`
	Version    int    `json:"version"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type AuditEvent struct {
	ID        string `json:"id"`
	RecordID  string `json:"record_id"`
	Action    string `json:"action"`
	Actor     string `json:"actor"`
	Detail    string `json:"detail"`
	CreatedAt string `json:"created_at"`
}

type Workflow struct {
	ID        string `json:"id"`
	RecordID  string `json:"record_id"`
	Stage     string `json:"stage"`
	Owner     string `json:"owner"`
	UpdatedAt string `json:"updated_at"`
}

type Attachment struct {
	ID       string `json:"id"`
	RecordID string `json:"record_id"`
	Name     string `json:"name"`
	Checksum string `json:"checksum"`
	Size     int    `json:"size"`
}

func NewRecord(id, title, summary, permission string) Record {
	return Record{ID: id, Title: title, Summary: summary, Permission: permission, Status: StatusDraft, Version: 1, CreatedAt: "deterministic", UpdatedAt: "deterministic"}
}

func (r Record) Validate() error {
	if r.ID == "" || r.Title == "" {
		return fmt.Errorf("id and title are required")
	}
	if r.Permission == "" {
		return fmt.Errorf("permission is required")
	}
	if r.Version < 1 {
		return fmt.Errorf("version must be positive")
	}
	return nil
}

func (r *Record) Transition(next Status) error {
	if r == nil {
		return fmt.Errorf("nil record")
	}
	if !CanTransition(r.Status, next) {
		return fmt.Errorf("invalid transition %s to %s", r.Status, next)
	}
	r.Status = next
	r.Version++
	r.UpdatedAt = "deterministic"
	return nil
}

func (r Record) IsVisible() bool { return r.Status != StatusArchived }
func (r Record) CanEdit() bool   { return r.Status == StatusDraft || r.Status == StatusReview }
func (r Record) Clone() Record   { return r }
