package audit

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

func ValidateRetention(events []domain.AuditEvent) error {
	for _, e := range events {
		if e.RecordID == "" || e.Action == "" {
			return fmt.Errorf("invalid audit event")
		}
	}
	return nil
}
func FilterActor(events []domain.AuditEvent, actor string) []domain.AuditEvent {
	out := []domain.AuditEvent{}
	for _, e := range events {
		if e.Actor == actor {
			out = append(out, e)
		}
	}
	return out
}
func ActionCounts(events []domain.AuditEvent) map[string]int {
	out := map[string]int{}
	for _, e := range events {
		out[e.Action]++
	}
	return out
}
