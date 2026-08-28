package store

import (
	"encoding/json"
	"example.com/cookproposal/internal/domain"
	"fmt"
)

func EncodeSnapshot(s Snapshot) ([]byte, error) { return json.Marshal(s) }
func DecodeSnapshot(data []byte) (Snapshot, error) {
	var s Snapshot
	err := json.Unmarshal(data, &s)
	return s, err
}
func (s *Store) ExportJSON() ([]byte, error) {
	snap, e := s.Snapshot()
	if e != nil {
		return nil, e
	}
	return EncodeSnapshot(snap)
}
func (s *Store) ImportSnapshot(snap Snapshot) error {
	if e := snap.Validate(); e != nil {
		return e
	}
	for _, r := range snap.Records {
		if e := s.SaveRecord(r); e != nil {
			return e
		}
	}
	for _, ev := range snap.Events {
		if e := s.SaveAudit(ev); e != nil {
			return e
		}
	}
	for _, w := range snap.Workflows {
		if e := s.SaveWorkflow(w); e != nil {
			return e
		}
	}
	for _, a := range snap.Attachments {
		if e := s.SaveAttachment(a); e != nil {
			return e
		}
	}
	return nil
}
func (s *Store) ReplaceRecord(r domain.Record) error {
	if e := s.RequireOpen(); e != nil {
		return e
	}
	existing, e := s.LoadRecord(r.ID)
	if e != nil {
		return e
	}
	if r.Version < existing.Version {
		return fmt.Errorf("stale version")
	}
	return s.SaveRecord(r)
}
