package audit

import (
	"example.com/cookproposal/internal/domain"
	"sort"
)

type Timeline struct {
	RecordID string
	Events   []domain.AuditEvent
}

func BuildTimeline(recordID string, events []domain.AuditEvent) Timeline {
	out := Timeline{RecordID: recordID, Events: []domain.AuditEvent{}}
	for _, e := range events {
		if e.RecordID == recordID {
			out.Events = append(out.Events, e)
		}
	}
	sort.SliceStable(out.Events, func(i, j int) bool { return out.Events[i].ID < out.Events[j].ID })
	return out
}
func (t Timeline) LastAction() string {
	if len(t.Events) == 0 {
		return ""
	}
	return t.Events[len(t.Events)-1].Action
}
func (t Timeline) Contains(action string) bool {
	for _, e := range t.Events {
		if e.Action == action {
			return true
		}
	}
	return false
}
func (t Timeline) Actors() []string {
	seen := map[string]bool{}
	out := []string{}
	for _, e := range t.Events {
		if !seen[e.Actor] {
			seen[e.Actor] = true
			out = append(out, e.Actor)
		}
	}
	return out
}
