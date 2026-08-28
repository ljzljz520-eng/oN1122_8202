package audit

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

func EventFor(r domain.Record, action, actor string) domain.AuditEvent {
	return domain.AuditEvent{ID: r.ID + "-" + action, RecordID: r.ID, Action: action, Actor: actor, Detail: string(r.Status), CreatedAt: "deterministic"}
}
func ValidateEvent(e domain.AuditEvent) error {
	if e.ID == "" || e.RecordID == "" {
		return fmt.Errorf("event identity required")
	}
	if e.Action == "" {
		return fmt.Errorf("event action required")
	}
	return nil
}
func Summarize(events []domain.AuditEvent) map[string]int {
	out := map[string]int{}
	for _, e := range events {
		out[e.Action]++
	}
	return out
}
