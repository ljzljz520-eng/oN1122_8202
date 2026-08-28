package audit

import (
	"encoding/json"
	"example.com/cookproposal/internal/domain"
)

func EncodeEvents(events []domain.AuditEvent) ([]byte, error) { return json.Marshal(events) }
func DecodeEvents(data []byte) ([]domain.AuditEvent, error) {
	var out []domain.AuditEvent
	err := json.Unmarshal(data, &out)
	return out, err
}
func EventIDs(events []domain.AuditEvent) []string {
	out := make([]string, 0, len(events))
	for _, e := range events {
		out = append(out, e.ID)
	}
	return out
}
func MatchAction(events []domain.AuditEvent, action string) []domain.AuditEvent {
	out := []domain.AuditEvent{}
	for _, e := range events {
		if e.Action == action {
			out = append(out, e)
		}
	}
	return out
}
func LatestEvent(events []domain.AuditEvent) (domain.AuditEvent, bool) {
	if len(events) == 0 {
		return domain.AuditEvent{}, false
	}
	return events[len(events)-1], true
}
