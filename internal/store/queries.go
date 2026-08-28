package store

import (
	"example.com/cookproposal/internal/domain"
	"fmt"
	"strings"
)

func (s *Store) FindByTitle(text string) ([]domain.Record, error) {
	records, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	needle := strings.ToLower(strings.TrimSpace(text))
	out := make([]domain.Record, 0)
	for _, r := range records {
		if needle == "" || strings.Contains(strings.ToLower(r.Title), needle) {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Store) FindByStatus(status domain.Status) ([]domain.Record, error) {
	records, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	out := make([]domain.Record, 0)
	for _, r := range records {
		if r.Status == status {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Store) FindByPermission(permission string) ([]domain.Record, error) {
	records, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	out := make([]domain.Record, 0)
	for _, r := range records {
		if r.Permission == permission {
			out = append(out, r)
		}
	}
	return out, nil
}
func (s *Store) Count() (int, error) { records, e := s.ListRecords(); return len(records), e }
func (s *Store) EnsureRecord(id string) (domain.Record, error) {
	r, e := s.LoadRecord(id)
	if e != nil {
		return domain.Record{}, fmt.Errorf("ensure %s: %w", id, e)
	}
	return r, nil
}
