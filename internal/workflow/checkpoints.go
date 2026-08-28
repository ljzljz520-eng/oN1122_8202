package workflow

import "example.com/cookproposal/internal/domain"

type Checkpoint struct {
	Stage    string
	Complete bool
	Message  string
}

func Check(r domain.Record) Checkpoint {
	switch r.Status {
	case domain.StatusDraft:
		return Checkpoint{Stage: "create", Message: "ready for review"}
	case domain.StatusReview:
		return Checkpoint{Stage: "review", Message: "awaiting confirmation"}
	case domain.StatusConfirmed:
		return Checkpoint{Stage: "confirm", Message: "ready to publish"}
	case domain.StatusPublished:
		return Checkpoint{Stage: "publish", Complete: true, Message: "published"}
	case domain.StatusArchived:
		return Checkpoint{Stage: "archive", Complete: true, Message: "archived"}
	default:
		return Checkpoint{Message: "unknown status"}
	}
}
func Complete(r domain.Record) bool { return Check(r).Complete }
func Checkpoints(records []domain.Record) []Checkpoint {
	out := make([]Checkpoint, 0, len(records))
	for _, r := range records {
		out = append(out, Check(r))
	}
	return out
}
