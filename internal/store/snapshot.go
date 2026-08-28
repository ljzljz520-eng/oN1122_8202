package store

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

type Snapshot struct {
	Records     []domain.Record
	Events      []domain.AuditEvent
	Workflows   []domain.Workflow
	Attachments []domain.Attachment
}

func (s *Store) Snapshot() (Snapshot, error) {
	if e := s.RequireOpen(); e != nil {
		return Snapshot{}, e
	}
	records, e := s.ListRecords()
	if e != nil {
		return Snapshot{}, e
	}
	events, e := s.ListAudits("")
	if e != nil {
		return Snapshot{}, e
	}
	snap := Snapshot{Records: records, Events: events}
	for _, r := range records {
		if w, err := s.LoadWorkflow("wf-" + r.ID); err == nil {
			snap.Workflows = append(snap.Workflows, w)
		}
		if a, err := s.LoadAttachment("att-" + r.ID); err == nil {
			snap.Attachments = append(snap.Attachments, a)
		}
	}
	return snap, nil
}
func (s Snapshot) Empty() bool {
	return len(s.Records) == 0 && len(s.Events) == 0 && len(s.Workflows) == 0 && len(s.Attachments) == 0
}
func (s Snapshot) Record(id string) (domain.Record, error) {
	for _, r := range s.Records {
		if r.ID == id {
			return r, nil
		}
	}
	return domain.Record{}, fmt.Errorf("record not in snapshot")
}
func (s Snapshot) Validate() error {
	for _, r := range s.Records {
		if e := r.Validate(); e != nil {
			return e
		}
	}
	return nil
}
func MergeSnapshots(a, b Snapshot) Snapshot {
	out := Snapshot{Records: append([]domain.Record{}, a.Records...), Events: append([]domain.AuditEvent{}, a.Events...), Workflows: append([]domain.Workflow{}, a.Workflows...), Attachments: append([]domain.Attachment{}, a.Attachments...)}
	out.Records = append(out.Records, b.Records...)
	out.Events = append(out.Events, b.Events...)
	out.Workflows = append(out.Workflows, b.Workflows...)
	out.Attachments = append(out.Attachments, b.Attachments...)
	return out
}
