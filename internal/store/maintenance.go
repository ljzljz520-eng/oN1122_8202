package store

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
)

type Health struct {
	Open    bool
	Records int
	Events  int
}

func (s *Store) Health() Health {
	if !s.Healthy() {
		return Health{}
	}
	records, _ := s.ListRecords()
	events, _ := s.ListAudits("")
	return Health{Open: true, Records: len(records), Events: len(events)}
}
func (s *Store) SaveRecordWithAudit(r domain.Record, e domain.AuditEvent) error {
	if err := s.SaveRecord(r); err != nil {
		return err
	}
	if e.ID == "" {
		e.ID = r.ID + "-change"
	}
	if e.RecordID == "" {
		e.RecordID = r.ID
	}
	return s.SaveAudit(e)
}
func (s *Store) ArchiveRecord(id, actor string) error {
	r, e := s.LoadRecord(id)
	if e != nil {
		return e
	}
	if e = r.Transition(domain.StatusArchived); e != nil {
		return e
	}
	return s.SaveRecordWithAudit(r, domain.AuditEvent{ID: id + "-archive", RecordID: id, Action: "archive", Actor: actor, Detail: "archived", CreatedAt: "deterministic"})
}
func (s *Store) RequireOpen() error {
	if !s.Healthy() {
		return fmt.Errorf("store closed")
	}
	return nil
}
